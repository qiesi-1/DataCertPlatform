package main

import (
	"data/blockchain"
	"data/db_mysql"
	_ "data/routers"
	"github.com/astaxie/beego"
)

func main() {
	//bc := blockchain.NewBlockChian()
	//fmt.Println("lastHash",bc.LastHash)
	//bc.SaveData([]byte("区块链C190604"))
	//blocks,err := bc.QueryAllBlocks()
	//if err!= nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	////blocks是一个切片
	//for index,block := range blocks {
	//	fmt.Printf("%d个区块%d高度%xhash%x\n",index,block.Height,block.PrevHash)
	//}
	//
	//return
	//先准备一条区块链
	blockchain.NewBlockChian()
	//链接数据库
	return
	db_mysql.Connect()

	//静态资源文件路径映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

