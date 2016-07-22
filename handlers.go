package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func NewRegisteration(w http.ResponseWriter, r *http.Request) {
	log.Println("New request from " + r.RemoteAddr + " " + r.RequestURI)
	output := JSON{ErrorCode: 0, Success: true, Result: nil}
	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		output.ErrorCode = 53
		output.Success = false
		output.Result = nil
	} else {
		var readuser ReadUser
		json.Unmarshal(jsonbody, &readuser)

		if !IsExists(readuser.Username) {
			token := NewUser(readuser.Username, readuser.Password)
			output.ErrorCode = 0
			output.Success = true
			output.Result = ReturnUser{Token: token}
		} else {
			output.ErrorCode = 51
			output.Success = false
			output.Result = ErrorValue{Detail: "UserExists"}
		}
	}
	json.NewEncoder(w).Encode(output)
}

func ValidateRegistration(w http.ResponseWriter, r *http.Request) {
	log.Println("New request from " + r.RemoteAddr + " " + r.RequestURI)
	output := JSON{ErrorCode: 0, Success: true, Result: nil}
	token := r.Header.Get("Authorization")
	userid := ValidateUser(token)
	if userid > 0 {
		NewLog(userid, 0, r.RequestURI, r.RemoteAddr)
		output.ErrorCode = 0
		output.Success = true
		output.Result = ErrorValue{Detail: "UserExists"}
	} else {
		output.ErrorCode = 52
		output.Success = false
		output.Result = ErrorValue{Detail: "UserNotExists"}
	}
	json.NewEncoder(w).Encode(output)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("New request from " + r.RemoteAddr + " " + r.RequestURI)
	output := JSON{ErrorCode: 0, Success: true, Result: nil}
	token := r.Header.Get("Authorization")
	userid := ValidateUser(token)
	if userid > 0 {
		NewLog(userid, 0, r.RequestURI, r.RemoteAddr)
		profile := GetProfileInfo(userid)
		if profile.Userid > 0 {
			output.ErrorCode = 0
			output.ErrorCode = 0
			output.Success = true
			output.Result = profile
		} else {
			output.ErrorCode = 53
			output.Success = false
			output.Result = ErrorValue{Detail: "NoProfileFounded"}
		}
	} else {
		output.ErrorCode = 52
		output.Success = false
		output.Result = ErrorValue{Detail: "UserNotExists"}
	}
	json.NewEncoder(w).Encode(output)
}

func SetProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("New request from " + r.RemoteAddr + " " + r.RequestURI)
	output := JSON{ErrorCode: 0, Success: true, Result: nil}
	token := r.Header.Get("Authorization")
	userid := ValidateUser(token)
	if userid > 0 {
		NewLog(userid, 0, r.RequestURI, r.RemoteAddr)
		jsonbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			output.ErrorCode = 53
			output.Success = false
			output.Result = ErrorValue{Detail: "InputJSONIsWrong"}
		} else {
			var readuser Profile
			json.Unmarshal(jsonbody, &readuser)
			readuser.Userid = userid
			SetProfileInfo(readuser)
			output.ErrorCode = 0
			output.Success = true
			output.Result = nil
		}
	} else {
		output.ErrorCode = 52
		output.Success = false
		output.Result = ErrorValue{Detail: "UserNotExists"}
	}
	json.NewEncoder(w).Encode(output)
}

func GetLastNotification(w http.ResponseWriter, r *http.Request) {
	log.Println("New request from " + r.RemoteAddr + " " + r.RequestURI)
	output := JSON{ErrorCode: 0, Success: true, Result: nil}
	token := r.Header.Get("Authorization")
	userid := ValidateUser(token)
	if userid > 0 {
		NewLog(userid, 0, r.RequestURI, r.RemoteAddr)
		notification := GetLastNotificationDB(userid)
		if notification.NotifyID > 0 {
			output.ErrorCode = 0
			output.Success = true
			output.Result = notification
			ReadLastNotificationDB(notification.NotifyID)
		} else {
			output.ErrorCode = 55
			output.Success = false
			output.Result = ErrorValue{Detail: "NoNotifyFounded"}
		}
	} else {
		output.ErrorCode = 52
		output.Success = false
		output.Result = ErrorValue{Detail: "UserNotExists"}
	}
	json.NewEncoder(w).Encode(output)
}

func SetDrink(w http.ResponseWriter, r *http.Request) {
	log.Println("New request from " + r.RemoteAddr + " " + r.RequestURI)
	output := JSON{ErrorCode: 0, Success: true, Result: nil}
	token := r.Header.Get("Authorization")
	userid := ValidateUser(token)
	if userid > 0 {
		NewLog(userid, 0, r.RequestURI, r.RemoteAddr)
		jsonbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			output.ErrorCode = 53
			output.Success = false
			output.Result = ErrorValue{Detail: "InputJSONIsWrong"}
		} else {
			var readuser Drink
			json.Unmarshal(jsonbody, &readuser)
			readuser.Userid = userid
			NewRecordDrink(readuser)
			output.ErrorCode = 0
			output.Success = true
			output.Result = nil
		}
	} else {
		output.ErrorCode = 52
		output.Success = false
		output.Result = ErrorValue{Detail: "UserNotExists"}
	}
	json.NewEncoder(w).Encode(output)
}
