package main

import (
	// "bufio"
	// "errors"
	// "log"

	// "encoding/json"
	// "errors"
	"fmt"
	// "os"
	// "strings"
	// "noteTaking/note/"
	"noteTaking/userInput"
	// "time"
)

func main(){

	title,content,err:=getNoteData()
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("title:",title)
	fmt.Println("content:",content)

}


func getNoteData()(title string,content string,err error){


	noteTitle,err:=userInput.GetUserINput("Note title:")
	// noteTitle,err:=usergetUserINput("Note title:")
	if err !=nil{

		return "","",err
		// fmt.Println(err)
	}
	fmt.Println(noteTitle)
	noteContent,err:=userInput.GetUserINput("Note content:")
	// noteContent,err:=getUserINput("Note content:")
	if err !=nil{
		return "","",err

		// fmt.Println(err)
	}
	fmt.Println(noteContent)

	return noteTitle,noteContent,nil
}
