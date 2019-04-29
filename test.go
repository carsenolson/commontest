package main

import (
	"os"
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
	test := Test.NewTest("test", 20)
	fmt.Println("Config initialized")
	fmt.Println("Config test and result paths: ", config.Test_path, config.Result_path)
	fmt.Println("Start creating 100 questions")
	for i := 0; i < 100; i++ {
		test.AddQuestion("test", []string{"test.png", "test.png"}, []string{"answer"+strconv.Itoa(i),
		"answer"+strconv.Itoa(i)}, []int{0,1})
	}
	fmt.Println("Created 100 questions")
	fmt.Println("Questions len is: ", len(test.Questions))
	fmt.Println("Creating test with 100 questions")
		fmt.Println("Created test with 100 questions")
	err = test.Save(config.Test_path, "test1.json")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println("Saved test 'test1.json' in tests path: ", config.Test_path)
	tests, err := Test.GetAllTests(config.Test_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("All tests:")
	for _, t := range tests {
		fmt.Println(t.Name())
	}
	test1, err := Test.NewTestFromFile(config.Test_path, "test1.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Parsed test from 'test1.json'")
	fmt.Println("tests questions length is: ", len(test1.Questions))
	fmt.Println("tests questions capacity is: ", cap(test1.Questions))
	test1.DeleteQuestion(34)
	test1.DeleteQuestion(34)
	test1.DeleteQuestion(len(test1.Questions)-1)
	fmt.Println("Deleted questions: 34, 34,", len(test1.Questions)-1)
	err = Test.DeleteTest(config.Test_path, "test1.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Deleted test 'test1.json'")
	fmt.Println("***TEST PASSED***")
}
