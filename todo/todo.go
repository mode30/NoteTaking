package todo

import (
	// "fmt"
	"encoding/json"
	"errors"
	// "errors"
	"fmt"
	// "log"
	"noteTaking/userInput"
	"os"
	// "strings"
	// "time"
)

type Todo struct{
	content string

}


func NewTodo( todo  string)(todoContent *Todo,err error){
	// todo.title,err=userInput.GetUserINput("enter title:")
	// if err !=nil{
	// 	fmt.Println(err)
	// 	return

	// }
	// todo,err=userInput.GetUserINput("enter content")
	// if err !=nil{
	// 	fmt.Println(err)
	// 	return

	// }
	// if todo.title ==""|| note.content==""{
	// 	return todo,errors.New("must have title and content filled")
	// }
	return &Todo{
		// title:todo.title,
		content:todo,
		// createdAt: time.Now(),

	},nil
}


func (todo *Todo)DisplayNote()(string,error){

	return fmt.Sprintf("information output:\n %s",todo.content),nil


}


func (todo *Todo)Save()(err error){
	// fileTodo:=strings.ReplaceAll(todo.content, "", "_")
	// fileTodo=strings.ToLower(fileTodo) + ".json"
	fileName:="todo.json"
	json,err:=json.Marshal(todo)
	if err!=nil{
		// fmt.Println(err)
		// log.Fatal(err)
		return err
	}
	return os.WriteFile(fileName, json, 0644)

}


func getTodoData()(todoMessage string,err error){
	userTodo,err:= userInput.GetUserINput("enter todo:")
	if err != nil{
		return "",errors.New("empty must return something")
	}
	return userTodo,nil
}
