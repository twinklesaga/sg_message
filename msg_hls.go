package sg_message


type MsgHlsPayload struct {
	AlbumId 	int		`json:"albumId"`
	TrackId		int 	`json:"trackId"`

	FtpPath		string  `json:"ftpPath"`
}


type MsgHls struct {
	Header MsgHeader 		`json:"header"`
	Payload MsgHlsPayload	`json:"payload"`
}
