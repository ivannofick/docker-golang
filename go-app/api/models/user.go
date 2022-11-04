package models

type User struct {
	Id          int64  `gorn:"primaryKey" json="id"`
	NamaLengkap string `gorn:"varchar(300)" json="nama_lengkap"`
	Username    string `gorn:"varchar(200)" json="username"`
	Password    string `gorn:"varchar(200)" json="password"`
}
