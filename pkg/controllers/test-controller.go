package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type UserResponse struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		Id        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
	Support struct {
		Url  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

func TestController(c echo.Context) error {
	logger := c.Logger()

	page := c.QueryParam("page")
	req, _ := http.NewRequest("GET", "https://reqres.in/api/users"+"?page="+page, nil)
	// Do send an HTTP request and
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Debug("error in send req: ", err)
	}
	if resp.Status != "200" {
		logger.Debug(resp.Status)
	}
	//Defer the closing of the body
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var data UserResponse
	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		logger.Debug(err)
	}
	return c.JSON(http.StatusOK, &data)

}
