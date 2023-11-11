package handlers

import "fmt"

func HandleError(err error) {
	if err != nil {
		fmt.Printf("[ERROR]: %v", err)
	}
}