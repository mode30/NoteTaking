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
	return &Todo{
		// title:todo.title,
		content:todo,
		// createdAt: time.Now(),

	},nil
}


func (todo *Todo)Display()(string){

	// return fmt.Sprintf("information output:\n %s",todo.content),nil
	result:=fmt.Sprintf("information output:\n %s",todo.content)
	return result


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
