package common

import "github.com/Chengxufeng1994/go-ddd/internal/domain/repository"

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var SUCCESS = 0

type PageResult struct {
	*repository.PaginationResult
	Rows interface{} `json:"rows"`
}

type PageResponse struct {
	Code int         `json:"code"`
	Data *PageResult `json:"data"`
	Msg  string      `json:"msg"`
}
