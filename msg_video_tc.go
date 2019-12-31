package sg_message

type MsgVideoTcPayload struct {
	RhId 			int        `json:"rhId"`
	RhName 			string   `json:"rhName"`
	Process 		string 		`json:"process"`
	Workspace 		string    `json:"workspace"`
	ManifestFile 	string     `json:"manifestFile"`
	Resource 		[]Resource `json:"resources"`
	AlbumId 		int         `json:"albumId"`
	TrackId 		int         `json:"trackId"`
	VideoId			int 		`json:"videoId"`
	VideoType		string 		`json:"videoType"`
	ResultDir 		string    `json:"resultDir"`
}

type MsgVideoTc struct {
	Header 		MsgHeader 			`json:"header"`
	Payload 	MsgVideoTcPayload	`json:"payload"`
}
