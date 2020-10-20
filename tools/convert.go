package tools

import (
	"bytes"
	"encoding/binary"
)

//将int64转化为字节切片
func Int64ToByte(num int64)([]byte,error) {
	buff := new(bytes.Buffer)//通过new实例化一个缓存区
	//buff.Write() 通过。。。向缓存区写入数据
	//buff.Bytes() 通过Bytes方法从缓存区获取数据
	
	//  大端位序排列 binary.BigEndian
	//  小端位序排列 binary.LittleEndian
	err := binary.Write(buff,binary.BigEndian,num)
	if err != nil{
		return nil,err
	}
	//从缓存区读取数据
	return buff.Bytes(),nil
}

//将字符串转化为字节
func StringToBytes(data string) []byte {
	return []byte(data)
}