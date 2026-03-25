package userInput

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"errors"
)

func GetUserINput(prompt string)(userQuery string,err error){

	fmt.Print(prompt)
	buffer:=bufio.NewReader(os.Stdin)
	bufferString,err:=buffer.ReadString('\n')
	if err !=nil{
		fmt.Println(err)
		return
		// log.Fatal(err)
	}
	userInput:=strings.TrimSpace(bufferString)
	if userInput == ""{
		return "",errors.New("user input cannot be empty")
	}
	return userInput,nil
}
