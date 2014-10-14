package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
			{
			    "clients": [
				{
				    "ip": "172.26.126.11",
				    "user": "root",
				    "pass": "root123"
				},
				{
				    "ip": "172.26.126.75",
				    "user": "root",
				    "pass": "root123"
				}
			    ]
			}
			`

	var m map[string][]interface{}

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		//var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%s: %s : %s\n", m.Name, m.Text, m.text2)
		//fmt.Printf("%s\n", m["clients"])

	}
	element := m["clients"]
	for _, abc := range element {

		fmt.Printf(" %s\n", abc)
	}

	type Response2 struct {
		IP   string   `json:"page"`
		User []string `json:"fruits"`
		//Pass string
	}
	//str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(jsonStream), &res)
	fmt.Println(res)
	fmt.Println(res.IP)
}
