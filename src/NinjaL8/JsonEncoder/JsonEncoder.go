/*Starting with this code, encode to JSON the []user sending the results to Stdout. Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})*
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Username1 string
	Password  string
}

func main() {
	u1 := User{
		Username1: "swapnil",
		Password:  "123456",
	}

	err := json.NewEncoder(os.Stdout).Encode(u1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u1)
}
