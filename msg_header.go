package sg_message

import (
	"github.com/satori/go.uuid"
	"strings"
	"time"
)



type MsgHeader struct {
	Gid 		string	`json:"gid"`
	Pid 		string	`json:"pid"`
	Mid 		string	`json:"mid"`
	Type 		string	`json:"type"`
	Version 	string	`json:"version"`
	From 		string	`json:"from"`
	To 			string 	`json:"to"`
	ReplyTo 	string	`json:"replyTo"`
	IssueDate 	int64	`json:"issueDate"`
	ExpiryDate 	int64	`json:"expiryDate"`
	Compressed 	string 	`json:"compressed"`
}




func GenID() string {
	id := uuid.NewV4()

	return strings.Replace(id.String() , "-" , "" , -1)
}

func GenNowDate() int64 {
	return  time.Now().Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func GenExiryDate(duration time.Duration) int64{
	return time.Now().Add(duration).Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
