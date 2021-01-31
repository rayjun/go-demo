package main

import (
	"fmt"
	"log"
	"os"

	"rayjun.cn/godemo/the_go_programming_language/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues: \n", result)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
