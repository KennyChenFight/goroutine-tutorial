package main

import "fmt"

type user struct {
	email string
	name string
	gender string
}

func main() {
	stringChan := make(chan string, 3)
	stringChan <- "hello world"
	//u := user{
	//	"kenny@example.com",
	//	"kenny",
	//	"male",
	//}
	userChan := make(chan user, 1)
	//userChan <- u
	test(userChan)
	fmt.Println(<- userChan)
	fmt.Println(userChan)
}

func test(userChan chan user) {
	u := user{
		"kenny@example.com",
		"kenny",
		"male",
	}
	userChan <- u
}
