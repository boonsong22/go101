package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://jsonplaceholder.typicode.com/users"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "76392230-336c-7ac9-5664-47f65304cfff")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var users []map[string]interface{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, user := range users {
		for key, val := range user {
			fmt.Println(key, val)
		}

	}
	// fmt.Printf("%+v", animals)
	// fmt.Println(res)
	// fmt.Println(string(body))

}
