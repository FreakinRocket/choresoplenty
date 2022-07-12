package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ChoreList struct {
	Chores []struct {
		Name  string `json:"name"`
		Last  string `json:"last"`
		Goals []struct {
			Name      string `json:"name"`
			Status    string `json:"status"`
			Value     string `json:"value"`
			Condition string `json:"condition"`
		} `json:"goals"`
	} `json:"chores"`
}

func main() {
	data := readChoresDB()
	fmt.Println(data)
}

func readChoresDB() ChoreList {
	file, _ := ioutil.ReadFile("choresDB.json")

	data := ChoreList{}

	_ = json.Unmarshal([]byte(file), &data)

	return data
}
