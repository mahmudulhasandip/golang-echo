package utils

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

type ResponseFormatStruct struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Lang    string      `json:"lang" default:"en"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseStruct struct {
	StatusCode int
	MessageEn  string `default:""`
	MessageBn  string `default:""`
}

func (rs ResponseStruct) ResponseFormat(data interface{}, lang string) ResponseFormatStruct {
	var message string
	if lang == "bn" {
		message = rs.MessageBn
	} else {
		message = rs.MessageEn
	}

	return ResponseFormatStruct{
		Code:    strconv.Itoa(rs.StatusCode),
		Message: message,
		Lang:    lang,
		Data:    data,
	}
}

func (rs ResponseStruct) WriteToResponse(c echo.Context, data interface{}, lang string) error {
	return c.JSON(rs.StatusCode, rs.ResponseFormat(data, lang))
}
