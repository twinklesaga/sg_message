package sg_message


type MsgArchiveMetaPayload struct {
	RhId 			int        	`json:"rhId"`
	RhName 			string   	`json:"rhName"`
	RhMetaCd		string		`json:"rhMetaCd"`
	MetaPath 		string 		`json:"metaPath"`
	MetaAlbums		[]int		`json:"metaAlbums"`
	ArchiveDate		string		`json:"ArchiveDate"`
}


type MsgArchiveMeta struct {
	Header 		MsgHeader 		`json:"header"`
	Payload 	MsgArchiveMetaPayload `json:"payload"`
}
