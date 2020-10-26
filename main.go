package main

import (
	"data/blockchain"
	"data/db_mysql"
	_ "data/routers"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	block0 := blockchain.CreateGenesisBlock()//创建创世区块
	block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte(""))
	fmt.Printf("block0的哈希:%x\n",block0.Hash)
	fmt.Printf("block1的prevhash:%x\n",block1.PrevHash)

	//序列化
	//将数据从内存中转化为可以持续存储在硬盘或在网络上传输的形式
	blockJson,_ := json.Marshal(block0)
	fmt.Println("json序列化后的block：",string(blockJson))
	blockXml,_ := xml.Marshal(block0)
	fmt.Println("json序列化后的block：",string(blockXml))
// 终止后续执行
	return

	//链接数据库
	db_mysql.Connect()
	//静态资源文件路径映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

