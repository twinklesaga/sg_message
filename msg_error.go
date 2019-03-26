package sg_message


type MsgErrorPayload struct {
	Message string 					`json:"message"`
	Ori interface{}					`json:"ori"`
}

type MsgError struct {
	Header 		MsgHeader 			`json:"header"`
	Payload 	MsgErrorPayload 	`json:"payload"`
}


