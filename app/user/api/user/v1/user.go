package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetUserReq struct {
	g.Meta `path:"/queryUser" tags:"queryUser" method:"get" summary:""`
	Id     int `json:"id"`
}
type GetUserRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
