package models

import (
	"bytes"
	"encoding/gob"
)

//该结构体用于定义数据保存的信息
type CertRecord struct {
	CertId []byte//认证号，md5值
	CertHash []byte//存正文件sha256值
	CertName string // 认证人的姓名
	Phone string //联系人手机号
	CertCard string //身份证
	FileName string//认证文件名
	FileSize int64 //文件大小
	CertTime int64 //认证时间
}


//序列化操作

func (c CertRecord)Serialize()([]byte,error)  {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(),err
}
//反序列化生成一个CertRecord结构体实例
func DeserializeCertRecord(data []byte) (*CertRecord,error){
	var certRecord *CertRecord
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&certRecord)
	return certRecord,err
}
