package models

import (
	"crypto/md5"
	"data/db_mysql"
	"encoding/hex"
	"fmt"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

//将用户信息保存至数据库
func (u User) AddUser() (int64, error) {
	//脱敏
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes) //把脱敏后的密码重新传入数据库
	rs, err := db_mysql.Db.Exec("insert into user(phone,password) values (?,?)",
		u.Phone, u.Password)
	//fmt.Println(u.Phone,u.Password,"输入框")
	if err != nil {
		//保存数据遇到错误
		return -1, err
		fmt.Println(err.Error())
	}

	id, err := rs.RowsAffected() //id 代表此次操作影响的行数，为int64类型
	if err != nil {
		//保存数据遇到错误
		return -1, err
	}
	return id, nil

}

func (u User) QueryUser() (*User, error) {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes)
	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
		u.Phone, u.Password)
	//fmt.Println(u.Phone,u.Password,"输出框")
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (u User) QueryUserByPhone()(*User,error){
	row := db_mysql.Db.QueryRow("select id from user where phone = ?",
		u.Phone)
	err := row.Scan(&u.Id)
	if err != nil{
		return nil,err
	}
	return &u,nil

}