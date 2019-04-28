package main

import (
	"fmt"
	"commontest/Test"
	"strconv"
)

func main() {
	questions := new([]Test.Question)
	for i := 0; i < 1000; i++ {
		*questions = append(*questions, *Test.NewQuestion("test", []string{"test.png", "test.png"},
		[]string{"answer"+strconv.Itoa(i), "answer"+strconv.Itoa(i)}, []int{0,1}))
	}
	for i, value := range *questions {
		fmt.Println("questions index and title", i, value)
	}
	fmt.Println("Questions len is: ", len(*questions))
	test := Test.NewTest("test", 20, *questions)
	err := test.Save("test1.json")
	if err != nil {
		fmt.Print(err)
	}
	test1, err := Test.NewTestFromFile("test1.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tests questions length is: ", len(test1.Questions))
	fmt.Println("tests questions capacity is: ", cap(test1.Questions))
	fmt.Println(test1.Questions[67])
	// It dosn't pass because of strange "out of range" panic	
	for i, value := range test1.Questions {
		if i > len(test1.Questions) {
			fmt.Println("Hm, something wrong")
		}
		fmt.Println("deleting question with index and value", i, value)
		fmt.Println("questions size: ", len(test1.Questions))
		test1.DeleteQuestion(i)
	}
	fmt.Println(test1.Questions, test1.Name, test1.Time)
}
