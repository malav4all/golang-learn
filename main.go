package main

import "fmt"

// someName := "Out Function"

func main() {
	// fmt.Println("Hello World!")

	//strings Declaration
	// var malav string = "Hello GoLang"
	// var malavTest = "Second Method"
	// var test string
	// fmt.Println(malav, malavTest, test)

	// test = "Reactjs"

	// fmt.Println(malav, malavTest, test)

	// malavFinal := "ShortHandway to declare string"
	// fmt.Println(malavFinal)

	// // Declare integer

	// var age int = 20
	// var ageTwo = 230
	// ageThree := 2000

	// fmt.Println(age, ageTwo, ageThree)

	//bits & Memory

	// var numOne int8 = 215

	// var numOne int16 = 215
	// var numTwo int8 = -128
	// var numThree uint8 = 25
	// var numFour uint16 = 65535

	//Float Value Declare

	// var scoreTwo float32 = 235.5656
	// var scoreThree float64 = 235.5656454546544645646465
	// scoreFor := 1.5
	// age := 25
	// name := "Malav"

	// fmt.Print("Malav, ")
	// fmt.Print("Naagar!")
	// fmt.Print("New Line \n")

	// fmt.Println("Malav Naagar")

	// fmt.Println("My age is", age, "My Name is", name)

	// //Printf (Formated String) %_  = Format Specifier

	// fmt.Printf("my age %v and my name is %v \n", age, name)
	// fmt.Printf("my age %q and my name is %q \n", age, name)

	// //to check age Type
	// fmt.Printf("agr is of type %T \n", age)
	// fmt.Printf("Your scored %f points! \n", 255.3)
	// fmt.Printf("Your scored %0.1f points! \n", 255.33)

	// // Sprintf (save Formatted String)

	// var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	// fmt.Println(str)

	// Array
	// var ages [3]int = [3]int{1, 2, 3}

	// Shorthand Array Declare
	// var ages = [3]int{1, 2, 3}

	// name := [4]string{"Malav", "Reacts", "Nodejs", "Test"}
	// fmt.Println(name, len(name))

	// Slice (use array Under The Hood)

	// var scores = []int{100, 200, 300}
	// scores[2] = 25
	// scores = append(scores, 400)
	// fmt.Println(scores, len(scores))

	// //slice ranges

	// rangeOne := name[1:3]
	// rangeTwo := name[2:]
	// rangeThree := name[:3]
	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// Standard Library

	// greeting := "Malav is a Best Developer"
	// fmt.Println(strings.Contains(greeting, "Malav")) // Return True or false String Search
	// fmt.Println(strings.ReplaceAll(greeting, "Developer", "SDE"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "Best"))

	// fmt.Println("Orginal Value", greeting)

	// loops

	// x := 0
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	names := []string{"Malav", "Test", "Reactjs", "SDE"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index, value := range names {
		fmt.Printf("The Postion at index %v is %v \n", index, value)
	}

}
