package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Title   string `json:"title"`
	Student string `json:"student"`
	Body    string `json:"body"`
}

/*
func main() {
	file, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println(err)
	}
	var data Data
	json.Unmarshal([]byte(file), &data)

	fmt.Printf("Title: %s, Student, %s, Body: %s", data.Title, data.Student, data.Body)

}
*/

func main() {
	file, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println(err)
	}
	var data Data
	json.Unmarshal([]byte(file), &data)

	m := Data{"Yeag", "Me?????", "yeah lol"}
	b, err := json.Marshal(m)

	ioutil.WriteFile("test.json", b, 0755)

	//file, err = ioutil.WriteFile("test.json")
}
