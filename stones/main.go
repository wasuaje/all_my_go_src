package main

import (
	"fmt"

)


func main() {
	entries := make([]int, 11 )
	for i := 0; i<=10; i++{
		entries[i] = i
	}
	for option:= range(entries){
		if option%7 == 0 || option%7 ==1{
			fmt.Println("Second")
		}else{
			fmt.Println("First")
		}
	}

}
