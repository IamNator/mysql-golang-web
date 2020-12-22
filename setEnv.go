package main

import (
	"fmt"
	"os"
)

func main(){
	as := os.Getenv("ACCESS_SECRET")
	if as == "" {
		err := os.Setenv("ACCESS_SECRET", "Our-web-app")
		if err != nil {
			fmt.Println("failed to set ACCESS_SECRET")
		}
	}
}
