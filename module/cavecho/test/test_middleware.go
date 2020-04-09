package main

import (
        "net/http"
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
        md "goAgent/module/cavecho"
	nd "goAgent"
)


func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func save1(c echo.Context) error {
        // Get name and email
        name := c.FormValue("name")
        email := c.FormValue("email")
        return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func save2(c echo.Context) error {
        // Get name and email
        name := c.FormValue("name")
        email := c.FormValue("email")
        return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func m1(bt uint64) {
        nd.Method_entry(bt, "m1")
        nd.Method_exit(bt, "m1")
}


func mainAdmin(c echo.Context)error{
	req := c.Request()

        ctx := req.Context()

        bt := ctx.Value("CavissonTx").(uint64)

        m1(bt)

	return c.String(http.StatusOK,"ID is coming")

}

func check1(c echo.Context)error{
      return c.String(http.StatusOK,"hey there id conding")

}

func ServerHeader(next echo.HandlerFunc)echo.HandlerFunc{
      return func (c echo.Context)error{
                  c.Response().Header().Set(echo.HeaderServer,"BlueBot/1.0")
                  return next(c)
	}
}


func main(){
	nd.Sdk_init()
	e:=echo.New()
	e.Use(ServerHeader)
	e.Use(md.Middleware())
	g:=e.Group("/admin")
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{

	Format:`[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` +"\n",

	}))
	g.Use(middleware.BasicAuth(func(username string,password string,c echo.Context)(bool,error){
	if username =="cavisson" && password =="cavisson"{
		return true,nil
	}
	return false,nil
	}))
	g.GET("/main",mainAdmin)
	e.POST("/cats",save)
        e.POST("/dog",save1)
        e.POST("/rat",save2)
        g.GET("/hero",check1)
	e.Start(":0000")
	nd.Sdk_free()
}
