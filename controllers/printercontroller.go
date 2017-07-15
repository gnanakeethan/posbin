package controllers

import (
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gnanakeethan/posbin/models"
)

// AuthenticationController operations for Authentication
type PrinterController struct {
	ActionController
}

func (c *PrinterController) URLMapping() {
	c.Mapping("Index", c.Index)
	//c.Mapping("AllBlock", c.AllBlock)
}

// @Title Get
// @Description get Bills by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PrinterController) Index() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUpdatedBill(id)
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	query["BillId"] = idStr
	sortby = append(sortby, "Id")
	order = append(order, "asc")
	p, err := models.GetAllSales(query, fields, sortby, order, 0, 100)

	logs.Info(p)
	logs.Info(v)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = p
	}
	c.ServeJSON()
}

/*
 $this->bill = Bill::find(session('bill_no'));
 $output = [];
 $output[] = ['text' => $this->getCompanyHeader(), 'em' => 1];
 $output[] = ['text' => $this->getBillHeader(), 'em' => 0];


 foreach ($this->bill->sales as $key => $sale) {
 if (!$sale->hide) {

 $output[] = [
 'text' => substr($this->getItemFormatted($key + 1, $sale->inventory->product->name, $sale->inventory->scale->first()->unit, $sale->inventory->batch_no, $sale->units, $sale->unit_price, $sale->amount, $sale->discount), 0, 32),
 'em'   => 1,
 ];
 $output[] = [
 'text' => substr($this->getItemFormatted($key + 1, $sale->inventory->product->name, $sale->inventory->scale->first()->unit, $sale->inventory->batch_no, $sale->units, $sale->unit_price, $sale->amount, $sale->discount), 32, 32),
 'em'   => 0,
 ];
 $output[] = [
 'text' => substr($this->getItemFormatted($key + 1, $sale->inventory->product->name, $sale->inventory->scale->first()->unit, $sale->inventory->batch_no, $sale->units, $sale->unit_price, $sale->amount, $sale->discount), 64, 32),
 'em'   => 0,
 ];

 $output[] = [
 'text' => substr($this->getItemFormatted($key + 1, $sale->inventory->product->name, $sale->inventory->scale->first()->unit, $sale->inventory->batch_no, $sale->units, $sale->unit_price, $sale->amount, $sale->discount), 96, 32),
 'em'   => 0,
 ];
 }
 };
 $output = array_merge($output, $this->getStatement());

 if ($this->bill->discount) {
 $output[] = ['text' => str_pad("DISCOUNTS", 32, "-", STR_PAD_BOTH) . "", 'em' => 1];
 }
 foreach ($this->bill->sales as $key => $sale) {
 if (!$sale->hide) {
 $output[] = [
 'text' => $this->getDiscounted($key + 1, $sale->inventory->product->name, $sale->inventory->scale->first()->unit, $sale->inventory->batch_no, $sale->units, $sale->unit_price, $sale->amount, $sale->discount),
 'em'   => 0,
 ];
 }
 }
 foreach ($output as $key => $out) {
 //            dd($output[0]['text']);
 if (isset($ouput[$key]['text']))
 $ouput[$key]['text'] = str_pad($out['text'], 32, ' ', STR_PAD_RIGHT);
 }
 $output[] = ['text' => $this->getThankYou(), 'em' => 1];
 $output[] = ['machine' => auth()->user()->terminal->machine, 'printer' => auth()->user()->terminal->printer];

 //        if ($this->bill->balance <= 0 && $this->bill->net_total > 0)
 $text='';
 foreach ($output as $out) {
 if (isset($out['text'])) {
 $text .= $out['text'];
 }
 }
 $textsplit = str_split($text,32);
 $text ="";
 $text .= str_pad("",48,"-",STR_PAD_BOTH)."\n";
 foreach($textsplit as $splt){
 $text .= str_pad($splt,48," ",STR_PAD_BOTH)."\n";
 }
 $output = [];
 $output[] = ['text' => $text];
 $output[] = ['machine' => auth()->user()->terminal->machine, 'printer' => auth()->user()->terminal->printer];

 return response()->json($output);
*/
