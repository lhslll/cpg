package main

import (
	_ "github.com/cpg/rcportal/routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	fmt.Println(beego.BConfig.WebConfig.ViewsPath)
	beego.Run()
}

