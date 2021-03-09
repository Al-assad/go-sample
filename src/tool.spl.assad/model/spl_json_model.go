package model

import "time"

// easyjson 测试模型

//easyjson:json
type Article struct {
	Title      string    `json:"title"`
	Auth       Author    `json:"auth"`
	CreateTime time.Time `json:"create_time"`
	Scope      int       `json:"scope"`
}

//easyjson:json
type Author struct {
	Name []string `json:"name"`
}
