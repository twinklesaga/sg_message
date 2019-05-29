package sg_message

type FtpFileInfo struct {
	Highlight 		bool			`json:"highlight,omitempty"`
	SvcType 		string			`json:"svcType,omitempty"`
	FileType 		string 			`json:"fileType,omitempty"`
	Bitrate 		int 			`json:"bitrate,omitempty"`
	SampleRate		int				`json:"sampleRate,omitempty"`
	StreamPath 		string 			`json:"streamPath,omitempty"`
	FileSize 		int64			`json:"fileSize,omitempty"`
	Duration 		int64			`json:"duration,omitempty"`
	Width 			int				`json:"width,omitempty"`
	Height 			int 			`json:"height,omitempty"`
	Url				string 			`json:"url,omitempty"`
	Type 			string 			`json:"type,omitempty"`
	SourceType		string 			`json:"sourceType,omitempty"`
}


type MsgDiPayload struct {
	RhId 			int           	`json:"rhId"`
	RhName 			string      	`json:"rhName"`
	MimeType		string        	`json:"mimeType"`
	ImageType 		string       	`json:"imageType"`
	Workspace 		string       	`json:"workspace"`
	ManifestFile 	string        	`json:"manifestFile"`
	Resource 		[]FtpFileInfo 	`json:"resources"`
	AlbumId 		int            	`json:"albumId"`
	TrackId 		int            	`json:"trackId"`
	ArtistId		int           	`json:"artistId"`
	ResultDir 		string       	`json:"resultDir"`
	Key				string     		`json:"key"`
	Eq				string      	`json:"eq"`
	TrackLoudness	string			`json:"trackLoudness"`
}

type MsgDI struct {
	Header MsgHeader 		`json:"header"`
	Payload MsgDiPayload	`json:"payload"`
}
