package sg_message


type MsgReTcPayload struct {
	AlbumId 		int         `json:"albumId"`
	TrackId 		int         `json:"trackId"`
	SourceFrom		string		`json:"sourceFrom"`
	Bucket			string 		`json:"bucket"`
	Source 			string 		`json:"source"`
	ResultDir 		string    	`json:"resultDir"`
}


type MsgReTC struct {
	Header 		MsgHeader 		`json:"header"`
	Payload 	MsgReTcPayload 	`json:"payload"`
}
