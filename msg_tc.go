package sg_message

import "time"

type MsgTcPayload struct {
	RhId 			int        `json:"rhId"`
	RhName 			string   `json:"rhName"`
	Workspace 		string    `json:"workspace"`
	ManifestFile 	string     `json:"manifestFile"`
	Resource 		[]Resource `json:"resources"`
	AlbumId 		int         `json:"albumId"`
	TrackId 		int         `json:"trackId"`
	ResultDir 		string    `json:"resultDir"`
	Policy			string   `json:"policy"`
}


type MsgTC struct {
	Header 		MsgHeader 		`json:"header"`
	Payload 	MsgTcPayload 	`json:"payload"`
}

func GenTCMessage(albumId int , trackId int , res Resource , resultDir string , policy string) MsgTC {
	Id := GenID()
	return MsgTC{
		Header:MsgHeader{
			Gid:        GenID(),
			Pid:        Id,
			Mid:        Id,
			Type:       "",
			Version:    "1.0.0",
			From:       "GEN",
			To:         "TC",
			ReplyTo:    "",
			IssueDate: 	GenNowDate(),
			ExpiryDate: GenExiryDate(time.Duration(time.Hour)),
			Compressed: "N",
		},
		Payload:MsgTcPayload{
			RhId:         0,
			RhName:       "FLO",
			Workspace:    "workspace",
			ManifestFile: "manifestFile",
			Resource:     []Resource{res},
			AlbumId:      albumId,
			TrackId:      trackId,
			ResultDir:    resultDir,
			Policy: policy,
		},
	}
}




