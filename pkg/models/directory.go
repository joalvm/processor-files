package models

type Directory struct {
	Id    string `json:"id"`
	Name  string `json:"Name"`
	Path  string `json:"path"`
	Index int    `json:"index"`
}
