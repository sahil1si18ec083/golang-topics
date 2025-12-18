package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"interface-module/note"
	"interface-module/todo"
)

type OutputTable interface {
	Display()
	Save() error
}

func saveData(data OutputTable) error {
	return data.Save()
}
func DisplayData(data OutputTable) {
	data.Display()
}
func printmethod_where_you_are_not_sure_about_type(val interface{}) {

	intval, ok := val.(int)

	if ok {
		fmt.Println(intval)

	}
	floatval, ok := val.(float64)
	if ok {
		fmt.Println(floatval)
	}

	switch val.(type) {
	case int:
		fmt.Println("integer")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("no valid ")
	}
	fmt.Println(val)
}
func main() {
	title, content := getNoteData()
	todotext := getTodoData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoNote, err := todo.New(todotext)

	if err != nil {
		fmt.Println(err)
		return
	}
	DisplayData(userNote)
	DisplayData(todoNote)
	err = saveData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = saveData(todoNote)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Saving the note succeeded!")
	printmethod_where_you_are_not_sure_about_type(89)
	printmethod_where_you_are_not_sure_about_type("hello")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}
func getTodoData() string {

	content := getUserInput("todo content:")

	return content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
