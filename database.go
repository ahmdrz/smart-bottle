package main

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func OpenDatabase() error {
	db, err = gorm.Open("mysql", "<username>:<passowrd>@/hooshnoosh?charset=utf8&parseTime=True&loc=Local")
	return err
}

func CloseDatabase() error {
	return db.Close()
}

func IsExists(user string) bool {
	err = OpenDatabase()
	if err != nil {
		panic(err)
	}
	userObject := DatabaseUser{}
	db.Table("users").Find(&userObject, "username = ?", user)
	err = CloseDatabase()
	if err != nil {
		panic(err)
	}

	return userObject.Userid > 0
}

func NewUser(user, pass string) string {
	token := GetMD5Hash(user + "mp&g" + pass)

	err = OpenDatabase()
	if err != nil {
		panic(err)
	}
	userObject := DatabaseUser{Username: user, Userpass: GetMD5Hash(pass), Token: token}
	db.Table("users").Create(&userObject)
	err = CloseDatabase()
	if err != nil {
		panic(err)
	}

	return token
}

func ValidateUser(token string) int {
	err = OpenDatabase()
	if err != nil {
		panic(err)
	}
	userObject := DatabaseUser{}
	db.Table("users").Find(&userObject, "token = ?", token)
	err = CloseDatabase()
	if err != nil {
		panic(err)
	}

	return userObject.Userid
}

func GetProfileInfo(userid int) Profile {
	err = OpenDatabase()
	if err != nil {
		panic(err)
	}
	userObject := Profile{}
	db.Table("profiles").Find(&userObject, "userid = ?", userid)
	err = CloseDatabase()
	if err != nil {
		panic(err)
	}

	return userObject
}

func SetProfileInfo(p Profile) {
	err = OpenDatabase()
	if err != nil {
		panic(err)
	}
	db.Table("profiles").Where("userid = ?", p.Userid).Delete(Profile{})
	db.Table("profiles").Create(&p)
	err = CloseDatabase()
	if err != nil {
		panic(err)
	}
}

func GetLastNotificationDB(userid int) Notify {
	err = OpenDatabase()
	if err != nil {
		panic(err)
	}
	userObject := Notify{}
	db.Table("notifies").Find(&userObject, "userid = ?", userid)
	err = CloseDatabase()
	if err != nil {
		panic(err)
	}
	return userObject
}

func ReadLastNotificationDB(notifyid int) {
	err = OpenDatabase()
	if err != nil {
		panic(err)
	}
	db.Table("notifies").Where("notify_id = ?", notifyid).Delete(Notify{})
	err = CloseDatabase()
	if err != nil {
		panic(err)
	}
}

func NewRecordDrink(d Drink) {
	err = OpenDatabase()
	if err != nil {
		panic(err)
	}
	db.Table("drinks").Create(&d)
	err = CloseDatabase()
	if err != nil {
		panic(err)
	}
}

func NewLog(userid, level int, detail, ip string) {
	err = OpenDatabase()
	if err != nil {
		panic(err)
	}
	db.Table("logs").Debug().Create(Log{Userid: userid, Level: level, Detail: detail, Ip: ip, Date: strconv.FormatInt(time.Now().UnixNano(), 10)})
	err = CloseDatabase()
	if err != nil {
		panic(err)
	}
}
