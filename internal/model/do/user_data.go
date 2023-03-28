package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserData is the golang structure of table user_data for DAO operations like Where/Data.
type UserData struct {
	g.Meta      `orm:"table:user_data, do:true"`
	Id          interface{} //
	Userid      interface{} //
	Useraddress interface{} //
	Parentid    interface{} //
}
