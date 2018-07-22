package greeting

import "fmt"

func Print(text string) {
	fmt.Println(text)
}

type Salutation struct {
	Name     string
	Greeting string
}

func (salutation *Salutation) Rename(newName string) {
	salutation.Name = newName
}

type Salutations []Salutation

func (salutations Salutations) Greet2(message string) {

	fmt.Println(message)
}

type Renameable interface {
	Rename(newName string)
}

func RenameToFrog(r Renameable) {
	r.Rename("Frog")
}

const (
	pi = 3.14
)

type Printer func(string)

func Greet(salutation Salutation, do func(string), do2 Printer) {

	// first, second := CreateMessage(salutation)
	// fmt.Println(first)
	// fmt.Println(second)

	// message, alternate := CreateMessage2(salutation)
	// fmt.Println(message)
	// fmt.Println(alternate)

	message3 := CreateMessage3(salutation.Name, salutation.Greeting, "another greeting")
	do(message3)
	do2(message3)
}

func CreateMessage(salutation Salutation) (string, string) {
	return salutation.Name + ": " + salutation.Greeting, "HEY " + salutation.Name
}

func CreateMessage2(salutation Salutation) (message string, alternate string) {
	message = salutation.Name + ": " + salutation.Greeting
	alternate = "HEY " + salutation.Name
	return
}

func CreateMessage3(name string, message ...string) string {
	//fmt.Println(len(message))
	return name + ": " + message[0]

}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func TypeSwitchTest(x interface{}) { //similiar to c# object
	switch x.(type) {
	case int:
		fmt.Println("INT")
	case string:
		fmt.Println("STRING")
	default:
		fmt.Println("UNKNOWN")
	}
}
