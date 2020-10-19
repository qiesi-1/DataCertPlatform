package tools

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

//对一个字符串进行hash
func MD5HashString(data string)string  {
	MD5Hash := md5.New()
	MD5Hash.Write([]byte(data))
	bytes := MD5Hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
func MD5HashReader(reader io.Reader)(string,error){
	md5Hash := md5.New()
	readerBytes,err := ioutil.ReadAll(reader)

	//fmt.Println("读取到的文件",readerBytes)

	if err!=nil {
		return "",err
	}
	md5Hash.Write(readerBytes)
	hashBytes :=md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}

//读取io流中的数据，并对数据进行hash计算，返回sha256值
func SHA256HashReader(reader io.Reader)(string,error)  {
	sha256Hash := sha256.New()
	readerBytes,err :=ioutil.ReadAll(reader)
	if err !=nil {
		return "",err
	}
	sha256Hash.Write(readerBytes)
	hashBytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}