package main

import (
	"fmt"
	"commontest/Test"
	"commontest/Config"
	"strconv"
)

func main() {
	config, err := Config.InitConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Config initialized")
	fmt.Println("Config test and result paths: ", config.Test_path, config.Result_path)
	fmt.Println("Start creating 1000 questions")
	questions := new([]Test.Question)
	for i := 0; i < 1000; i++ {
		*questions = append(*questions, *Test.NewQuestion("test", []string{"test.png", "test.png"},
		[]string{"answer"+strconv.Itoa(i), "answer"+strconv.Itoa(i)}, []int{0,1}))
	}
	fmt.Println("Created 1000 questions")
	fmt.Println("Questions len is: ", len(*questions))
	fmt.Println("Creating test with 1000 questions")
	test := Test.NewTest("test", 20, *questions)
	fmt.Println("Created test with 1000 questions")
	err = test.Save(config.Test_path+"/"+"test1.json")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Saved test 'test1.json' in tests path: ", config.Test_path)
	test1, err := Test.NewTestFromFile(config.Test_path+"/"+"test1.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Parsed test from 'test1.json'")
	fmt.Println("tests questions length is: ", len(test1.Questions))
	fmt.Println("tests questions capacity is: ", cap(test1.Questions))
	test1.DeleteQuestion(134)
	test1.DeleteQuestion(356)
	test1.DeleteQuestion(len(test1.Questions)-1)
	fmt.Println("Deleted questions: 134, 356, ", len(test1.Questions)-1)
	Test.DeleteTest(config.Test_path+"/"+"test1.json")
	fmt.Println("Deleted test 'test1.json'")
	fmt.Println("TEST PASSED")
}
