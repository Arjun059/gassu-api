package main

import "fmt"

type Author struct {
	Name string
}

type Blog struct {
	Title   string
	Content string
	Author  Author
}

// return result strung S
// write to std output  P
// write to File or io written F

func hi(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

func main() {

	// blog := Blog{

	// 	Title:   "Nice Title",
	// 	Content: "Nice Content",
	// 	Author: Author{
	// 		Name: "Tomi",
	// 	},
	// }

	hiCaller := hi("Arjun")

	fmt.Println(hiCaller)

	// fmt.Printf("blog title %s content {%s} Author %+v blog Author Name %s", blog.Title, blog.Content, blog.Author, blog.Author.Name)

}
