package models

type Error struct {
	ErrorCode int `json:"err_code"`
	ErrorMessage string `json:"err_msg"`
}