package sg_message

type MsgPubPayload struct {
	RhId 			int			`json:"rhId"`
	RhName 			string		`json:"rhName"`
	MimeType		string 		`json:"mimeType"`
	Workspace 		string		`json:"workspace"`
	ManifestFile 	string		`json:"manifestFile"`
	Resource 		[]Resource	`json:"resources"`
	AlbumId 		int			`json:"albumId"`
	TrackId 		int 		`json:"trackId"`
	ResultDir 		string		`json:"resultDir"`
	Key				string 		`json:"key"`
	Eq 				string 		`json:"eq"`
	TrackLoudness	string		`json:"trackLoudness"`
	Duration		int			`json:"duration"`
	SourceType		string 		`json:"sourceType"`
	OriginFile 		string 		`json:"originFile"`
}


type MsgPub struct {
	Header 		MsgHeader 		`json:"header"`
	Payload 	MsgPubPayload 	`json:"payload"`
}


