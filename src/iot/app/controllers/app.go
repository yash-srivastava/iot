package controllers

import (
	"github.com/revel/revel"
	"iot/lib/tcp_server"
	"iot/lib/job_worker"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	go job_worker.Init()
	go tcp_server.Start_tcp_server()
	return c.Render()
}
