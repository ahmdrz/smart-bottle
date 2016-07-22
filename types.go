package main

type JSON struct {
	Success   bool        `json:"success"`
	ErrorCode int8        `json:"errorcode"`
	Result    interface{} `json:"result"`
}

type ReadUser struct {
	Username string
	Password string
}

type ReturnUser struct {
	Token string `json:"token"`
}

type ErrorValue struct {
	Detail string `json:"errorname"`
}

type DatabaseUser struct {
	Userid   int    `json="userid"`
	Username string `json="username"`
	Userpass string `json="userpass"`
	Token    string `json="token"`
}

type Profile struct {
	Age    int `json="age"`
	Height int `json="height"`
	Weight int `json="weight"`
	Userid int `json="userid"`
	IsMale int `json="is_male"`
}

type Notify struct {
	Level    int    `json="level"`
	Content  string `json="content"`
	Title    string `json="title"`
	NotifyID int    `json="id"`
}

type Drink struct {
	Drinkid int    `json="drink_id"`
	Userid  int    `json="user_id"`
	Value   int    `json="value"`
	Date    string `json="date"`
}

type Log struct {
	Logid  int    `json="log_id"`
	Userid int    `json="user_id"`
	Level  int    `json="level"`
	Detail string `json="detail"`
	Ip     string `json="ip"`
	Date   string `json="date"`
}
