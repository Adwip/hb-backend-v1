package model

type WebResponse struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"msg"`
	Data    interface{} `json:"data"`
}

type RepoResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     interface{} `json:"msg"`
}
