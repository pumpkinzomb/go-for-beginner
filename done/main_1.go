package main

import (
	"fmt"
	"go-for-beginner/dictionary"
)


func main(){
	mydict := dictionary.Dictionary{"book1": "rice"}
	fmt.Println(mydict)
	value, err := mydict.Get("book1")
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println("book1: ", value)
	err2 := mydict.Edit("book1", "vegitable")
	if(err2 != nil){
		fmt.Println(err2)
	}
	found, err := mydict.Get("book1")
	fmt.Println("book1: ", found)
}
