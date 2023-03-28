package entity

// UserData is the golang structure for table user_data
type UserData struct {
	Id          int    `json:"id"          description:""`
	Userid      string `json:"userid"      description:""`
	Useraddress string `json:"useraddress" description:""`
	Parentid    int    `json:"parentid"    description:""`
}
