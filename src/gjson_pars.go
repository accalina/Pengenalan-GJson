package src

import (
	"fmt"
	"io/ioutil"

	"github.com/imroc/req/v2"
	"github.com/tidwall/gjson"
)

func Parse_from_string() {
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

	data := gjson.Parse(mystr)
	fmt.Println(data.Get("success"))
	fmt.Println(data.Get("status.code"))

	for _, member := range data.Get("members").Array() {
		fmt.Println(member.Get("name"))
		fmt.Println(member.Get("gender").String() + "\n")
	}

}

func Parse_from_file() {
	raw, _ := ioutil.ReadFile("member.json")
	data := gjson.Parse(string(raw))
	fmt.Println(data.Get("success"))
	fmt.Println(data.Get("status.code"))

	for _, member := range data.Get("members").Array() {
		fmt.Println(member.Get("name"))
		fmt.Println(member.Get("gender").String() + "\n")
	}

}

func Parse_from_url() {
	r, _ := req.Get("http://localhost:3000/data")
	content, _ := r.ToString()

	data := gjson.Parse(content)
	fmt.Println(data.Get("success"))
	fmt.Println(data.Get("status.code"))

	for _, member := range data.Get("members").Array() {
		fmt.Println(member.Get("name"))
		fmt.Println(member.Get("gender").String() + "\n")
	}

}
