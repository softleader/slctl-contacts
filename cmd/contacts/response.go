package main

import (
	"fmt"
	"github.com/gosuri/uitable"
)

type contacts struct {
	Header []interface{}   `json:"header"`
	Datas  [][]interface{} `json:"datas"`
}

func (c *contacts) horizontalTable() (table *uitable.Table) {
	table = uitable.New()
	table.AddRow(c.Header...)
	for _, data := range c.Datas {
		table.AddRow(data...)
	}
	return
}

func (c *contacts) verticalTable() (table *uitable.Table) {
	table = uitable.New()
	for _, data := range c.Datas {
		for i, header := range c.Header {
			table.AddRow(fmt.Sprintf("%s:", header), data[i])
		}
		table.AddRow("") // blank
	}
	return
}
