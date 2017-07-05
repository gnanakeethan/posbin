package migrator

import "fmt"

type Migrate struct {
	TableName      string
	Engine         string
	Charset        string
	ModifyType     string
	Columns        []*Column
	Indexes        []*Index
	Primary        []*Column
	Uniques        []*Unique
	Foreigns       []*Foreign
	Renames        []*RenameColumn
	RemoveColumns  []*Column
	RemoveIndexes  []*Index
	RemoveUniques  []*Unique
	RemoveForeigns []*Foreign
}

type Index struct {
	Name string
}
type Foreign struct {
	Column        *Column
	ForeignTable  string
	ForeignColumn string
}
type Unique struct {
	Definition string
	Columns    []*Column
}

type Column struct {
	Name     string
	Inc      string
	Null     string
	Default  string
	Unsign   string
	DataType string
	remove   bool
	Modify   bool
}
type RenameColumn struct {
	OldName     string
	OldNull     string
	OldDefault  string
	OldUnsign   string
	OldDataType string
	NewName     string
	Column
}

/*




Column Methods start
*/

func (m *Migrate) NewCol(name string) *Column {
	col := &Column{Name: name}
	m.AddColumns(col)
	return col
}
func (m *Migrate) PriCol(name string) *Column {
	col := &Column{Name: name}
	m.AddColumns(col)
	m.AddPrimary(col)
	return col
}
func (m *Migrate) UniCol(uni, name string) *Column {
	col := &Column{Name: name}
	m.AddColumns(col)

	uniqueOriginal := &Unique{}

	for _, unique := range m.Uniques {
		if unique.Definition == uni {
			unique.AddColumns(col)
			uniqueOriginal = unique
		}
	}
	if uniqueOriginal.Definition == "" {
		unique := &Unique{Definition: uni}
		unique.AddColumns(col)
		m.AddUnique(unique)
	}

	return col
}

func (c *Column) Remove() {
	c.remove = true
}

func (c *Column) SetAuto(inc bool) *Column {
	if inc {
		c.Inc = "auto_increment"
	}
	return c
}
func (c *Column) SetNullable(null bool) *Column {
	if null {
		c.Null = "DEFAULT NULL"

	} else {
		c.Null = "NOT NULL"
	}
	return c
}
func (c *Column) SetDefault(def string) *Column {
	c.Default = def
	return c
}
func (c *Column) SetUnsigned(unsign bool) *Column {
	if unsign {
		c.Unsign = "UNSIGNED"
	}
	return c
}
func (c *Column) SetDataType(dataType string) *Column {
	c.DataType = dataType
	return c
}

func (c *RenameColumn) SetOldNullable(null bool) *RenameColumn {
	if null {
		c.OldNull = "DEFAULT NULL"

	} else {
		c.OldNull = "NOT NULL"
	}
	return c
}
func (c *RenameColumn) SetOldDefault(def string) *RenameColumn {
	c.OldDefault = def
	return c
}
func (c *RenameColumn) SetOldUnsigned(unsign bool) *RenameColumn {
	if unsign {
		c.OldUnsign = "UNSIGNED"
	}
	return c
}

func (c *RenameColumn) SetOldDataType(dataType string) *RenameColumn {
	c.OldDataType = dataType
	return c
}

func (c *Column) SetPrimary(m *Migrate) *Column {
	m.Primary = append(m.Primary, c)
	return c
}

func (m *Migrate) NewUnique(name string) (unique *Unique) {
	unique = &Unique{Definition: name}
	return
}
func (unique *Unique) AddColumns(columns ...*Column) *Unique {
	for _, column := range columns {
		unique.Columns = append(unique.Columns, column)
	}
	return unique
}

/*


Colum Methods end













*/

/*




 Migration Methods


*/

func (migration *Migrate) AddColumns(columns ...*Column) *Migrate {
	for _, column := range columns {
		migration.Columns = append(migration.Columns, column)
	}
	return migration
}
func (migration *Migrate) AddPrimary(primary *Column) *Migrate {
	migration.Primary = append(migration.Primary, primary)
	return migration
}

func (migration *Migrate) AddUnique(unique *Unique) *Migrate {
	migration.Uniques = append(migration.Uniques, unique)
	return migration
}

func (migration *Migrate) AddForeign(foreign *Foreign) *Migrate {
	migration.Foreigns = append(migration.Foreigns, foreign)
	return migration
}
func (migration *Migrate) AddIndex(index *Index) *Migrate {
	migration.Indexes = append(migration.Indexes, index)
	return migration
}

func (migration *Migrate) RenameColumn(from, to string) *RenameColumn {
	rename := &RenameColumn{OldName: from, NewName: to}
	migration.Renames = append(migration.Renames, rename)
	return rename
}

func (migration *Migrate) GetSQL() (sql string) {
	sql = ""
	switch migration.ModifyType {
	case "create":
		{
			sql += fmt.Sprintf("CREATE TABLE `%s` (", migration.TableName)
			for index, column := range migration.Columns {
				sql += fmt.Sprintf("\n `%s` %s %s %s %s %s", column.Name, column.DataType, column.Unsign, column.Null, column.Inc, column.Default)
				if len(migration.Columns) > index+1 {
					sql += ","
				}
			}

			if len(migration.Primary) > 0 {
				sql += fmt.Sprintf(",\n PRIMARY KEY( ")
			}
			for index, column := range migration.Primary {
				sql += fmt.Sprintf(" `%s`", column.Name)
				if len(migration.Primary) > index+1 {
					sql += ","
				}

			}
			if len(migration.Primary) > 0 {
				sql += fmt.Sprintf(")")
			}

			for _, unique := range migration.Uniques {
				sql += fmt.Sprintf(",\n UNIQUE KEY `%s`( ", unique.Definition)
				for index, column := range unique.Columns {
					sql += fmt.Sprintf(" `%s`", column.Name)
					if len(unique.Columns) > index+1 {
						sql += ","
					}
				}
				sql += fmt.Sprintf(")")
			}
			sql += fmt.Sprintf(")ENGINE=%s DEFAULT CHARSET=%s;", migration.Engine, migration.Charset)
			break
		}
	case "alter":
		{
			sql += fmt.Sprintf("ALTER TABLE `%s` ", migration.TableName)
			for index, column := range migration.Columns {
				sql += fmt.Sprintf("\n ADD `%s` %s %s %s %s %s", column.Name, column.DataType, column.Unsign, column.Null, column.Inc, column.Default)
				if len(migration.Columns) > index+1 {
					sql += ","
				}
			}
			for index, column := range migration.Renames {
				sql += fmt.Sprintf(",\n CHANGE COLUMN `%s` `%s` %s %s %s %s %s", column.OldName, column.NewName, column.DataType, column.Unsign, column.Null, column.Inc, column.Default)
				if len(migration.Renames) > index+1 {
					sql += ","
				}
			}

			sql += ";"

			break
		}
	case "reverse":
		{

			sql += fmt.Sprintf("ALTER TABLE `%s`", migration.TableName)
			for index, column := range migration.Columns {
				sql += fmt.Sprintf("\n DROP COLUMN `%s`", column.Name)
				if len(migration.Columns) > index+1 {
					sql += ","
				}
			}

			if len(migration.Primary) > 0 {
				sql += fmt.Sprintf(",\n DROP PRIMARY KEY")
			}

			for index, unique := range migration.Uniques {
				sql += fmt.Sprintf(",\n DROP KEY `%s`", unique.Definition)
				if len(migration.Uniques) > index+1 {
					sql += ","
				}

			}
			for index, column := range migration.Renames {
				sql += fmt.Sprintf(",\n CHANGE COLUMN `%s` `%s` %s %s %s %s", column.NewName, column.OldName, column.OldDataType, column.OldUnsign, column.OldNull, column.OldDefault)
				if len(migration.Renames) > index+1 {
					sql += ","
				}
			}
			sql += ";"
		}
	case "delete":
		{
			sql += fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", migration.TableName)
		}
	}

	return
}
