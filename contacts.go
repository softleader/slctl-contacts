package main

type contacts struct {
	Header []interface{}   `json:"header"`
	Datas  [][]interface{} `json:"datas"`
}
