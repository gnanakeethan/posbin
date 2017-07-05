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
}
type RenameColumn struct {
	OldName string
	NewName string
}

/*




Column Methods start
 */
//func (c *Column) SetName(name string) *Column {
//	c.Name = name
//	return c
//}
//func (c *Column) SetAuto(inc bool) *Column {
//	if inc {
//		c.Inc = "auto_increment"
//	}
//	return c
//}
//func (c *Column) SetNullable(null bool) *Column {
//	if null {
//		c.Null = "NULLABLE"
//
//	} else {
//		c.Null = "NOT NULL"
//	}
//	return c
//}
//func (c *Column) SetDefault(def string) *Column {
//	c.Default = def
//	return c
//}
//func (c *Column) SetUnsigned(unsign bool) *Column {
//	if unsign {
//		c.Unsign = "UNSIGNED"
//	}
//	return c
//}
//func (c *Column) SetDataType(dataType string) *Column {
//	c.DataType = dataType
//	return c
//}
/*


Colum Methods end













 */

/*




 Migration Methods


*/
func (unique *Unique) AddColumns(columns ...*Column) *Unique {
	for _, column := range columns {
		unique.Columns = append(unique.Columns, column)
	}
	return unique
}

func NewMigrate() (migration Migrate) {
	migration = Migrate{}
	return migration

}
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

func (migration *Migrate) RenameColumn(column *RenameColumn) *Migrate {
	migration.Renames = append(migration.Renames, column)
	return migration
}

func (migration *Migrate) GetSQL() (sql string) {
	sql = ""
	switch migration.ModifyType {
	case "create":
		{
			sql += fmt.Sprintf("CREATE TABLE `%s` (", migration.TableName)
			for index, column := range migration.Columns {
				sql += fmt.Sprintf("\n `%s` %s %s %s %s", column.Name, column.DataType,column.Unsign,column.Null,column.Inc)
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
			sql += fmt.Sprintf("ALTER TABLE `%s` (", migration.TableName)
			break
		}
	case "delete":
		{
			sql += fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", migration.TableName)
		}
	}

	return
}
