package main

import (
	"data/blockchain"
	"data/db_mysql"
	_ "data/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//1.创世区块

	bc := blockchain.NewBlockChian()
	fmt.Printf("创世区块的hash:%x\n",bc.LastHash)
	return
	//链接数据库
	db_mysql.Connect()
	//静态资源文件路径映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

