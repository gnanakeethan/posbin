package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:BillsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:BillsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:BillsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:BillsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:BillsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:BillsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:BillsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:BillsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:BillsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:BillsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:CustomersController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:CustomersController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:CustomersController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:CustomersController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:CustomersController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:CustomersController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:CustomersController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:CustomersController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:CustomersController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:CustomersController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountablesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountablesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountablesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountablesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountablesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountablesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountablesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountablesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountablesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountablesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:DiscountsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoriesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoriesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoriesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoriesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoriesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoriesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoriesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoriesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoriesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoriesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoryScaleController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoryScaleController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoryScaleController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoryScaleController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoryScaleController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoryScaleController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoryScaleController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoryScaleController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoryScaleController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:InventoryScaleController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:JobsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:JobsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:JobsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:JobsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:JobsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:JobsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:JobsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:JobsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:JobsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:JobsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PermissionsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PermissionsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PermissionsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PermissionsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PermissionsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductProductController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductProductController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductProductController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductProductController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductProductController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductProductController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductProductController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductProductController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductProductController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductProductController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"],
		beego.ControllerComments{
			"GetCount",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ProductsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PurchasesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PurchasesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PurchasesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PurchasesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PurchasesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PurchasesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PurchasesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PurchasesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PurchasesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:PurchasesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:RolesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:RolesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:RolesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:RolesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:RolesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:SalesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:SalesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:SalesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:SalesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:SalesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:SalesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:SalesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:SalesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:SalesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:SalesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ScalesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ScalesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ScalesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ScalesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ScalesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ScalesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ScalesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ScalesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ScalesController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:ScalesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:StocksController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:StocksController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:StocksController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:StocksController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:StocksController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:StocksController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:StocksController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:StocksController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:StocksController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:StocksController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:TerminalsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:TerminalsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:TerminalsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:TerminalsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:TerminalsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:TerminalsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:TerminalsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:TerminalsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:TerminalsController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:TerminalsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:UsersController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:UsersController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:UsersController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:UsersController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/gnanakeethan/pospo5/controllers:UsersController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
