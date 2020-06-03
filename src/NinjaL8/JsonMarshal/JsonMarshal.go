/*Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.
 */

package main

import (
	"encoding/json"
	"fmt"
)

type User struct{
	Username1 string
	Password string
}

func main(){
	u1:= User{
		Username1: "swapnil",
		Password: "123456",
	}

	fmt.Println(u1)

	b,err:=json.Marshal(u1)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(b))
}