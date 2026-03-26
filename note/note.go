package note

import (
	// "fmt"
	"encoding/json"
	"errors"
	"fmt"
	// "log"
	"noteTaking/userInput"
	"os"
	"strings"
	"time"
)

type NoteTaking struct{
	title string
	content string
	createdAt time.Time

}


func newNote(note *NoteTaking)(noteContent *NoteTaking,err error){
	note.title,err=userInput.GetUserINput("enter title:")
	if err !=nil{
		fmt.Println(err)
		return

	}
	note.content,err=userInput.GetUserINput("enter content")
	if err !=nil{
		fmt.Println(err)
		return

	}
	if note.title ==""|| note.content==""{
		return note,errors.New("must have title and content filled")
	}
	return &NoteTaking{
		title:note.title,
		content:note.content,
		createdAt: time.Now(),

	},nil
}


func (note *NoteTaking)displayNote()(string,error){

	return fmt.Sprintf("information output:\n %s,%s,%s",note.title,note.content,note.createdAt.Format(time.RFC3339)),nil


}


func (note *NoteTaking)Save()(err error){
	fileNote:=strings.ReplaceAll(note.title, "", "_")
	fileNote=strings.ToLower(fileNote) + ".json"
	json,err:=json.Marshal(note)
	if err!=nil{
		// fmt.Println(err)
		// log.Fatal(err)
		return err
	}
	return os.WriteFile(fileNote, json, 0644)

}
