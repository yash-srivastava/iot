package controllers

import (
	"github.com/revel/revel"
	"iot/lib/sender"
	"iot/lib/utils"
)

type App struct {
	*revel.Controller
}


func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Send_3000() revel.Result {

	response := sender.Response{Success: false, Message:"Something went wrong"}

	params := sender.Packet_3000{}
	params.SguId = utils.ToUint64(c.Params.Form.Get("sgu_id"))
	params.ScuId = utils.ToUint64(c.Params.Form.Get("scu_id"))
	params.Pwm = utils.ToInt(c.Params.Form.Get("pwm"))
	params.Op1 = utils.ToInt(c.Params.Form.Get("op1"))
	params.Op2 = utils.ToInt(c.Params.Form.Get("op2"))
	params.Op3 = utils.ToInt(c.Params.Form.Get("op3"))
	params.Op4 = utils.ToInt(c.Params.Form.Get("op4"))
	params.Retry = utils.ToInt(c.Params.Form.Get("retry"))
	params.RetryDelay = utils.ToInt(c.Params.Form.Get("retry_delay"))

	response = sender.HandlePacket(0x3000, params)
	return c.RenderJSON(response)
}