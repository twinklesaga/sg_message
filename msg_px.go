package sg_message


type MsgPXPayload struct {
	RhId 			int				`json:"rhId"`
	RhName 			string			`json:"rhName"`
	OrgWorkspace 	string 			`json:"orgWorkspace"`
	Workspace 		string 			`json:"workspace"`
	ManifestFile 	string 			`json:"manifestFile"`
	Resources 		[]PXResource  	`json:"resources"`
	ResultDir 		string 			`json:"resultDir"`
}

type PXResource struct {
	Path 			string 			`json:"path"`
	LastModified 	int64 			`json:"lastModified"`
}

type MsgPX struct {
	Header 			MsgHeader 		`json:"header"`
	Payload 		MsgPXPayload 	`json:"payload"`
}
