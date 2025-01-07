package models

type Directory struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	RealName string `json:"real_name"`
	Index    int    `json:"index"`
	DirIndex int    `json:"dir_index"`
}
