package entity

type UserInfo struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Sex    string `json:"sex"`
}
