package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/imroc/req/v2"
)

type Response struct {
	Success bool `json:"success"`
	Status  struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"status"`
	Members []struct {
		Name   string `json:"name"`
		Gender string `json:"gender"`
	}
}

func Default_parse_json() {
	mystr := `{
		"success": true,
		"status": {
			"code": 200,
			"msg": "ok"
		},
		"members": [
			{"name": "will", "gender": "male"},
			{"name": "lily", "gender": "female"}
		]}`
	data := Response{}
	json.Unmarshal([]byte(mystr), &data)

	fmt.Println(data.Success)
	fmt.Println(data.Status.Code)

	for _, member := range data.Members {
		fmt.Println(member.Name)
		fmt.Println(member.Gender + "\n")
	}
}

func Default_parse_file() {
	raw, _ := ioutil.ReadFile("member.json")

	data := Response{}
	json.Unmarshal([]byte(raw), &data)

	fmt.Println(data.Success)
	fmt.Println(data.Status.Code)

	for _, member := range data.Members {
		fmt.Println(member.Name)
		fmt.Println(member.Gender + "\n")
	}
}

func Default_parse_url() {
	raw, _ := req.Get("http://localhost:3000/data")
	content, _ := raw.ToBytes()

	data := Response{}
	json.Unmarshal([]byte(content), &data)

	fmt.Println(data.Success)
	fmt.Println(data.Status.Code)

	for _, member := range data.Members {
		fmt.Println(member.Name)
		fmt.Println(member.Gender + "\n")
	}

}
