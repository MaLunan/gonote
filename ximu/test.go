package main

import "fmt"

func main() {
	key := ""
	for {
		fmt.Print("请选择: ")

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("ssss")
		}
	}
}
