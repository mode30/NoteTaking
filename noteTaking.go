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
	"noteTaking/todo"
	// "time"
)

type saver interface{
	Save() error

}
// type displayer interface{
// 	Display()string
// }

type outputtable interface{
	// Display()string
	// Save()error
	saver
	Display()
}
func main(){

	printSomething("hello")

	todo1,err:=userInput.GetUserINput("enter todo:")
	if err !=nil{
		fmt.Println(err)
		return
	}

	//create new todo list
	todoContent,err:=todo.NewTodo(todo1)
	if err != nil{
		fmt.Println(err)
		return
	}

	// Display todo list
	todoResult:=todoContent.Display()
	fmt.Println("Display:",todoResult)
	// outputtable(todoContent.Display())
	// if err !=nil{
	// 	fmt.Println(err)
	// 	return
	// }

	//save in json file
	 err=saveData(todoContent)
	if err !=nil{
		fmt.Println(err)
		return
	}




	//Note
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

func saveData(data saver)(err error){

	err=data.Save()
	if err !=nil{
		fmt.Println("cannot save data")
		return err
	}

	fmt.Println("todo list saved")
	return nil
}

// func displayData(data displayer){
// 	data.Display()
// }

func DisplaySave(data outputtable){
	data.Display()
	saveData(data )
	// data.Save()
}


func printSomething(value any){
	fmt.Println(value)

}
