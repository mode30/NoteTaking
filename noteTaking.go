package main

import (
	// "bufio"
	// "errors"
	// "log"

	// "encoding/json"
	// "errors"
	"fmt"
	"go/types"
	// "os"
	// "strings"
	// "noteTaking/note/"
	"noteTaking/todo"
	"noteTaking/userInput"
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
	returnedValue:=returnGenericType(1,2)
	fmt.Println("return value:",returnedValue)

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


func returnGenericType[T  string|float64|int ](a,b T)(T){
	return a+b

}
func addNumber(a,b any)(result any,err error){
	aIntValue,isAInta:=a.(int)
	bIntValue,isBIntb:=a.(int)

	if isAInta && isBIntb{
		return aIntValue + bIntValue,nil
	}


	aFloat64Value,isAFloat64a:=a.(int)
	bFloat64Value,isBFloat64b:=a.(int)

	if isAFloat64a && isBFloat64b{
		return aFloat64Value + bFloat64Value,nil
	}
	return
}
func printSomething(value any){
	float64Value,ok:=value.(float64)
	if !ok{
		intValue,ok:=value.(int)
		if !ok{

		}
		fmt.Printf("int value:%d",intValue)
	}
	fmt.Printf("float64 %0.2f:",float64Value+1)
	// switch value.(type){
		// case int:
		// fmt.Println(value)
		// case float64:
		// fmt.Println("float64",value)
		// // fmt.Printf("%.2f",float64(value) *2)
	// }
	// fmt.Println(value)

}
