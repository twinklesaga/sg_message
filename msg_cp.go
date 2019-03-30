package sg_message

import "time"

type MsgCpPayload struct {
	RhId 			int        	`json:"rhId"`
	RhName 			string   	`json:"rhName"`
	Workspace 		string    	`json:"workspace"`
	ManifestFile 	string     	`json:"manifestFile"`
	Resource 		[]Resource 	`json:"resources"`
	AlbumId 		int         `json:"albumId"`
	TrackId 		int         `json:"trackId"`
	ResultDir 		string    	`json:"resultDir"`
	Next			string 		`json:"next"`
}


type MsgCP struct {
	Header 		MsgHeader 		`json:"header"`
	Payload 	MsgCpPayload 	`json:"payload"`
}

func GenCpMessage(albumId int , trackId int , res Resource , resultDir string , next string) MsgCP {
	Id := GenID()

	return MsgCP{
		Header:MsgHeader{
			Gid:        GenID(),
			Pid:        Id,
			Mid:        Id,
			Type:       "",
			Version:    "1.0.0",
			From:       "GEN",
			To:         "CP",
			ReplyTo:    "",
			IssueDate: 	GenNowDate(),
			ExpiryDate: GenExiryDate(time.Duration(time.Hour)),
			Compressed: "N",
		},
		Payload:MsgCpPayload{
			RhId:         0,
			RhName:       "FLO",
			Workspace:    "workspace",
			ManifestFile: "manifestFile",
			Resource:     []Resource{res},
			AlbumId:      albumId,
			TrackId:      trackId,
			ResultDir:    resultDir,
			Next: next,
		},
	}
}


