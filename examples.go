package main

import (
	"fmt"

	"./greeting"
)

func mainexamples() {

	//var s = greeting.Salutation{"Anne", "Hello"}
	//fmt.Println(s.greeting)
	// prefix := ""
	// switch s.Name {
	// case "Bob":
	// 	prefix = "Mr"
	// case "Sally", "Anne":
	// 	prefix = "Mrs"
	// }

	// switch {
	// case s.Name == "Bob":
	// 	prefix = "Mr"
	// case len(s.Name) == 4:
	// 	prefix = "Ms"
	// }

	//greeting.Greet(s, greeting.Print, greeting.CreatePrintFunction("!!!"))

	//fmt.Println("Hello " + prefix + s.Name)

	//greeting.TypeSwitchTest(123)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("times")
	// }

	// j := 4
	// times := 10

	// for j < times { //while
	// 	fmt.Println("times")
	// 	j++
	// }

	// i := 0 // do while
	// for {
	// 	if i%2 == 0 {
	// 		i++
	// 		continue
	// 	}
	// 	if i > 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// slice := []greeting.Salutation{
	// 	{"Bob", "Yo"},
	// 	{"Sally", "Hi"},
	// }

	// for _, s := range slice {
	// 	fmt.Println(s.Name)
	// }

	// var prefixMap map[string]string
	// prefixMap = make(map[string]string)

	// prefixMap["bob"] = "mr"
	// prefixMap["sally"] = "mr"

	prefixMap := map[string]string{
		"bob":   "mr",
		"sally": "mrs",
	}

	delete(prefixMap, "sally")

	//fmt.Println(prefixMap)

	if value, exists := prefixMap["bob"]; exists { // multiple returns
		fmt.Println(value)
	}

	//Slice is reference type, array is fixed size and value type
	// var s []int
	// s = make([]int, 3, 10)
	// s[0] = 1
	// s[1] = 2
	// s[2] = 3
	s := []int{1, 2, 3}

	s = s[0:2]

	//s[2] = 5
	s = append(s, 6)
	fmt.Println(s)

	s = append(s, s...)
	fmt.Println(s)

	slice := greeting.Salutations{
		{"Bob", "Yo"},
		{"Sally", "Hi"},
	}

	slice.Greet2(slice[0].Name)

	slice[0].Rename("Tom")

	done := make(chan bool)

	go func() {
		slice.Greet2(slice[0].Name)
		done <- true
	}()

	greeting.RenameToFrog(&slice[0]) //address of this so passing pointer in

	//time.Sleep(100 * time.Millisecond)
	<-done // blocks till it can read a value from done

	fmt.Println("done")

	done2 := make(chan bool, 2)
	go func() {
		slice.Greet2(slice[0].Name)
		done2 <- true
		done2 <- true
		fmt.Println("done2")
	}()

	greeting.RenameToFrog(&slice[0]) //address of this so passing pointer in

	//time.Sleep(100 * time.Millisecond)
	<-done // blocks till it can read a value from done

	fmt.Println("done")

	//c := make(chan int)
	//

}
