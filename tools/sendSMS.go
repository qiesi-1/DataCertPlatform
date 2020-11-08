package tools

import (
	"encoding/json"
	"github.com/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/beego"
)
//该函数用于发送一条短信
/*
	phone 电话号码
	code 发送的验证码
	template 模板
*/

type SmsCode struct {
	Code string`json:"code"`
}

type SmsResult struct {
	BizId respon
}

func SendSMS(phone string,code string,templateType string) (error) {
	config := beego.AppConfig
	//获取配置文件
	accessKey :=config.String("sms_access_key")
	accessKeySecret :=config.String("sms_access_secret")
	client,err :=dysmsapi.NewClientWithAccessKey("cn-hangzhou",accessKey,accessKeySecret)
	if err!= nil {
		return err
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = phone//指定要发送的目标手机
	request.SignName = "线上餐厅"
	request.TemplateCode = templateType

	smsbytes,_ := json.Marshal(SmsCode)
	request.TemplateParam =string(smsbytes)//指定发送验证码
	response,err := client.SendSms(request)
	if err!=nil {
		return err
	}

	SmsResult{
		BizId:
			}


}

func RandCode(width int) string {
	var smsLogin models.SmsLogin
}