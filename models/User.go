package models

import (
	"data/db_mysql"
	"data/tools"
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
	 u.Password = tools.MD5HashString(u.Password)//把脱敏的密码的MD5值重新赋值为密码进行储存
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
	//将脱敏的密码的MD5值重新赋值为密码进行储存
	u.Password = tools.MD5HashString(u.Password)

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
	var user User
	err := row.Scan(&user.Id)
	if err != nil{
		return nil,err
	}
	return &user,nil

}