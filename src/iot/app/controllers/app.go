package controllers

import (
	"github.com/revel/revel"
	"iot/lib/sender"
	"iot/conf"
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
	packet := sender.Packet_3000{}
	err := c.Params.BindJSON(&packet)
	if err==nil{
		response = sender.HandlePacket(0x3000, packet)
	}
	return c.RenderJSON(response)
}

func (c App) Send_8000() revel.Result {

	response := sender.Response{Success: false, Message:"Something went wrong"}
	packet := sender.Packet_8000{}
	err := c.Params.BindJSON(&packet)
	if err==nil{
		response = sender.HandlePacket(0x8000, packet)
	}
	return c.RenderJSON(response)
}

func (c App) IsSguConnected() revel.Result {
	response := sender.Response{Success: false, Message:"Something went wrong"}

	params := make(map[string]uint64)
	err := c.Params.BindJSON(&params)
	if err==nil{
		response.Data = conf.SGU_TCP_CONNECTION.Has(utils.ToStr(params["sguid"]))
		response.Success = true
		response.Message = ""
	}
	return c.RenderJSON(response)
}

func (c App) GetConnectedScus() revel.Result {
	response := sender.Response{Success: false, Message:"Something went wrong"}

	params := make(map[string]uint64)
	err := c.Params.BindJSON(&params)
	if err==nil{

		response.Data,_ = conf.SGU_SCU_LIST.Get(utils.ToStr(params["sguid"]))
		response.Success = true
		response.Message = ""
	}
	return c.RenderJSON(response)
}
