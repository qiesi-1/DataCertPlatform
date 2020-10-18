package controllers

import (
	"data/models"
	"data/tools"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)
//该结构体用于处理文件上传
type UploadFileController struct {
	beego.Controller
}
// 该post 方法用于处理用户在客户端提交的文件
func (u *UploadFileController) Post() {

	phone := u.Ctx.Request.PostFormValue("phone")
	title := u.Ctx.Request.PostFormValue("upload_title")
	fmt.Println("电子数据标签", title)
	file, header, err := u.GetFile("file")
	if err != nil {

		return
	}
	defer file.Close()

	//使用io包中方法保存文件
	saveFilePath := "static/upload/" + header.Filename
	_ ,err = tools.SaveFile(saveFilePath,file)
	if err != nil {
		u.Ctx.WriteString("抱歉，数据文件认证失败，请重试")
		return
	}
	//计算文件SHA256值
	tools.SHA256HashReader(file)

	//先查询用户ID
	user1,err := models.User{Phone:phone}.QueryUserByPhone()
	if err != nil{
		u.Ctx.WriteString("抱歉，电子数据认证失败，请稍后再试：")
		return
	}

	//把上传的文件作为记录保存到数据库中
	// ① 计算MD5 值
	md5String,err :=tools.MD5HashReader(file)
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败")
		return
	}
	recode := models.UploadRecord{
		UserId:user1.Id,
		FileName:header.Filename,
		FileSize:header.Size,
		FileCert:md5String,
		FileTitle:title,
		CertTime:time.Now().Unix(),
	}
	_,err = recode.SaveRecord()
	if err!=nil {
		u.Ctx.WriteString("抱歉，电子数据认证保存失败，请稍后再试‘")
		return
	}
	//上传文件保存到数据库 数据列表
	recodes,err := models.QueryRecordsByUserId(user1.Id)
	if err != nil{
		u.Ctx.WriteString("抱歉，获取电子数据列表失败")
		return
	}
	u.Ctx.WriteString("恭喜，已接受到上传文件")
	u.Data["Records"] = recodes
	u.TplName = "list_code.html"
}

//
//func (u *UploadFileController) Post()  {
//
//	title := u.Ctx.Request.PostFormValue("upload_title")
//	//
//	file,header,err := u.GetFile("file")
//	if err!= nil{
//		//解析客户端提交文件出错
//		u.Ctx.WriteString("抱歉，文件解析失败，请重试")
//		return
//	}
//	defer file.Close()
//	fmt.Println("自定义的标题",title)
//	//获得到上传文件
//	fmt.Println("上传的文件名：",header.Filename)
//	fileNameSlice := strings.Split(header.Filename,".")
//	fileType := fileNameSlice[1]
//	fmt.Println()
//	if fileType != "jpg"|| fileType != "png"{
//		u.Ctx.WriteString("文件上传格式错误，请上传符合格式的文件")
//		return
//	}
//	//isJpg := strings.HasSuffix(header.Filename,"jpg")
//	//isPng := strings.HasSuffix(header.Filename,"png")
//	//if !isJpg &&!isPng {
//	//	u.Ctx.WriteString("文件上传格式错误，请上传符合格式的文件")
//	//	return
//	//}
//	fmt.Println("上传文件大小",header.Size)
//	config := beego.AppConfig
//	fileSize,err := config.Int64("file_size")
//	if header.Size / 1024 > fileSize {
//		u.Ctx.WriteString("文件过大,请更换文件")
//		return
//	}
//	fmt.Println(file)
//	//fromFile：文件，
//	//toFile： 要保存的文件
//	saveDir := "sattic/upload"
//	f,err :=os.Open(saveDir)
//	if err!= nil{
//		err = os.Mkdir(saveDir,777)
//		if err != nil{
//			u.Ctx.WriteString("抱歉，文件认证出错，重试")
//			return
//		}
//	}
//	fmt.Println(f.Name())
//
//	savaName := saveDir+"/"+ header.Filename
//	u.SaveToFile("file",savaName)
//	if err!= nil{
//		u.Ctx.WriteString("抱歉，文件认证失败")
//		return
//	}
//	u.Ctx.WriteString("获取到上传。")
//	fmt.Println(savaName)
//}
