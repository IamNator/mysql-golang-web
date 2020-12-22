package main

import (
	"fmt"
	"os"
)

func main(){
	as := os.Getenv("ACCESS_SECRET")
	//fmt.Println(string(as))
	if as != "Our-web-app" {
		err := os.Setenv("ACCESS_SECRET", "Our-web-app")
		if err != nil {
			fmt.Println("failed to set ACCESS_SECRET")
		}
		fmt.Println("ACCESS_SECRET Set!")
	} else {
		fmt.Println("ACCESS_SECRET Already set")
	}
}
