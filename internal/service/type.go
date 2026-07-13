package service

import (
	"fmt"
)

type Error struct {
	Code int `json:"code"`
	Msg string `json:"message"`
	Err error  `json:"error"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s : %v",e.Msg,e.Err)
}

func (e Error) Unwrap() error {
	return e.Err
}

type Success struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}
