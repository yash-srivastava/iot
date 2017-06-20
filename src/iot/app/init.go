package app

import (
	"github.com/revel/revel"
	"iot/lib/tcp_server"
	"iot/lib/job_worker"
	"time"
	"iot/conf"
	"github.com/bsphere/le_go"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string

)

func SetTimeFormat(){
	loc, _ := time.LoadLocation("Asia/Kolkata")
	time.Local = loc
}

func InitForConn(){
	conf.Init()
}

func SetLogger()  {
	token := revel.Config.StringDefault("logentries","")
	if token!=""{
		le, err := le_go.Connect(revel.Config.StringDefault("logentries",""))
		if err != nil {
			panic(err)
		}
		revel.INFO.SetOutput(le)
		revel.WARN.SetOutput(le)
		revel.ERROR.SetOutput(le)
	}

}

func StartupScript() {
	SetTimeFormat()
	InitForConn()
	SetLogger()
	go job_worker.Init()
	go tcp_server.Start_tcp_server()
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	revel.OnAppStart(StartupScript)
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}