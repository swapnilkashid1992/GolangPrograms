/*Starting with this code, unmarshal the JSON into a Go data structur*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username1 string `json:"Username1"`
	Password  string `json:"Password"`
}

func main() {
	jsonData := `{"Username1":"swapnil","Password":"123456"}`
	var u User
	err := json.Unmarshal([]byte(jsonData), &u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
