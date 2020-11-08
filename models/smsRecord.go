package models

import "data/db_mysql"

type SmsRecord struct {
	BizId string
	Phone string
	Code string
	Status string
	Message string
	Timestamp int64
}

func(s SmsRecord) SaveSmsRecord(){
	rs,err := db_mysql.Db.Exec("insert into sms_record(biz_id,phone,code,status,message,timestamp),value (?,?,?,?,?)"
	s.BizId,s.Phone,s.Code,s.Message,s.Status,s.Timestamp)
}
