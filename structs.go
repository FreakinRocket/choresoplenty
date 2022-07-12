package main

type Chore struct {
	Name  string `json:"name"`
	Last  string `json:"last"`
	Goals []struct {
		Name      string `json:"name"`
		Status    string `json:"status"`
		Value     string `json:"value"`
		Condition string `json:"condition"`
	} `json:"goals"`
}
