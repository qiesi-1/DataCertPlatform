package controllers

import (
	"data/blockchain"
	"data/models"
	"fmt"
	"github.com/beego"
)

type CertDetailController struct {
	beego.Controller
}

//该get方法用于接受处理get请求
func (c *CertDetailController)Get()  {
	//解析和接受前端页面传递的CertId
	cert_id := c.GetString("cert_id")
	//查询区块数据
	block,err :=blockchain.CHAIN.QueryBloockByCertId(cert_id)
	if err != nil{
		c.Ctx.WriteString("抱歉，查询链上数据失败，请重试！")
		return
	}
	if block == nil{//遍历整条区块链未查到数据
		c.Ctx.WriteString("抱歉，未查到链上数据")
		return
	}
	fmt.Println("查询到的高度",block.Height)

	//反序列化
	certRecord,err :=models.DeserializeCertRecord(block.Data)
	//结果体
	c.Data["CertRecord"] = certRecord
	// 跳转证书详情
	c.TplName ="cert_detail.html"
}