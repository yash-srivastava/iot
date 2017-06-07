package controllers

import (
	"github.com/revel/revel"
	"iot/lib/sender"
	"encoding/json"
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

	params := c.Params.JSON
	err := json.Unmarshal(params, &packet)
	if err==nil{
		response = sender.HandlePacket(0x3000, packet)
	}

	return c.RenderJSON(response)
}