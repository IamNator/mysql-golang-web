package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(os.UserHomeDir())
	as := os.Getenv("ACCESS_SECRET")
	fmt.Println("ACCESS_SECRET Already set")
	if as == "" {
		err := os.Setenv("ACCESS_SECRET", "Our-web-app")
		if err != nil {
			fmt.Println("failed to set ACCESS_SECRET")
		}
		fmt.Println("ACCESS_SECRET Set!")
	}
}
