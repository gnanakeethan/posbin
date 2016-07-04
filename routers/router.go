// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/gnanakeethan/pospo5/controllers"

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

		beego.NSNamespace("/stocks",
			beego.NSInclude(
				&controllers.StocksController{},
			),
		),

		beego.NSNamespace("/terminals",
			beego.NSInclude(
				&controllers.TerminalsController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
