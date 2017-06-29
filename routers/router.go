// Package routers implement the routing functionality of the system
//
// @APIVersion 0.2.1-ab
// @Title posbin
// @Description REST API
// @Contact kee@pos.run
// @TermsOfServiceUrl http://www.pos.run/tos
// @License Proprietary
// @LicenseUrl http://www.pos.run/license
package routers

import (
	"github.com/gnanakeethan/posbin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/bills",
			beego.NSInclude(
				&controllers.BillsController{},
			),
		),

		beego.NSNamespace("/customers",
			beego.NSInclude(
				&controllers.CustomersController{},
			),
		),

		beego.NSNamespace("/discountables",
			beego.NSInclude(
				&controllers.DiscountablesController{},
			),
		),

		beego.NSNamespace("/discounts",
			beego.NSInclude(
				&controllers.DiscountsController{},
			),
		),

		beego.NSNamespace("/inventories",
			beego.NSInclude(
				&controllers.InventoriesController{},
			),
		),

		beego.NSNamespace("/inventory_scale",
			beego.NSInclude(
				&controllers.InventoryScaleController{},
			),
		),

		beego.NSNamespace("/jobs",
			beego.NSInclude(
				&controllers.JobsController{},
			),
		),

		beego.NSNamespace("/permissions",
			beego.NSInclude(
				&controllers.PermissionsController{},
			),
		),

		beego.NSNamespace("/product_product",
			beego.NSInclude(
				&controllers.ProductProductController{},
			),
		),

		beego.NSNamespace("/products",
			beego.NSInclude(
				&controllers.ProductsController{},
			),
		),

		beego.NSNamespace("/purchases",
			beego.NSInclude(
				&controllers.PurchasesController{},
			),
		),

		beego.NSNamespace("/roles",
			beego.NSInclude(
				&controllers.RolesController{},
			),
		),

		beego.NSNamespace("/sales",
			beego.NSInclude(
				&controllers.SalesController{},
			),
		),

		beego.NSNamespace("/scales",
			beego.NSInclude(
				&controllers.ScalesController{},
			),
		),

		beego.NSNamespace("/stock_flows",
			beego.NSInclude(
				&controllers.StockFlowsController{},
			),
		),
		beego.NSNamespace("/stocks",
			beego.NSInclude(
				&controllers.StocksController{},
			),
		),
		beego.NSNamespace("/stores",
			beego.NSInclude(
				&controllers.StoresController{},
			),
		),

		beego.NSNamespace("/terminals",
			beego.NSInclude(
				&controllers.TerminalsController{},
			),
		),

		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthenticationController{},
			),
		),
		beego.NSNamespace("/printer",
			beego.NSInclude(
				&controllers.PrinterController{},
			),
		),
		//beego.AutoRouter(&controllers.AuthenticationController{}),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/", &controllers.HomeController{}, "*:Index")
}
