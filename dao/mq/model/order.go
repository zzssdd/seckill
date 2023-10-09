package model

type Order struct {
	Id        int64 `json:"id"`
	Uid       int64 `json:"uid"`
	Pid       int   `json:"pid"`
	TimeStamp int64 `json:"timeStamp"`
}
