package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
)

const ERROR_MESSAGE = "FAILED_TO_GET_DATA"
const GET_ALL_URL = "http://pokeapi.co/api/v2/pokedex/kanto/"

func main() {
	e := echo.New()
	e.GET("/", getAllData)
	e.Start(":1111")
}

func getAllData(c echo.Context) error {
	data, err := callGetAllApi(GET_ALL_URL)
	if err != nil {
		return c.JSON(http.StatusNotFound, ERROR_MESSAGE)
	}
	return c.JSON(http.StatusOK, data)
}

func callGetAllApi(url string) (map[string]interface{}, error) {
	r, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(responseData, &result); err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
