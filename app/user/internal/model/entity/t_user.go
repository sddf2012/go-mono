// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TUser is the golang structure for table t_user.
type TUser struct {
	Id     uint   `json:"id"     orm:"id"     description:"user id"`     // user id
	Name   string `json:"name"   orm:"name"   description:"user name"`   // user name
	Status int    `json:"status" orm:"status" description:"user status"` // user status
	Age    uint   `json:"age"    orm:"age"    description:"user age"`    // user age
}
