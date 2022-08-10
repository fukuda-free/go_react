package main

import (
	_ "sample/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlConn,err := beego.AppConfig.String("sqlconn")
	if err != nil {
		beeLogger.Log.Fatal("%s", err)
	}
	orm.RegisterDataBase("default", "mysql", sqlConn)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

