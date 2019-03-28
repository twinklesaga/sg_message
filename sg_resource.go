package sg_message

type Resource struct {
	Path 		string			`json:"path"`
	Bitrate 	int 			`json:"bitrate"`
	Target 		string 			`json:"target"`
	UploadType	string			`json:"uploadType"`
	Type 		string 			`json:"type"`
}
