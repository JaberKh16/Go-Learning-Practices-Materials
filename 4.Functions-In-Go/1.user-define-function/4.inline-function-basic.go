package main

import "fmt"

func main(){
	limit := 10
	for idx :=1; idx <=limit; idx++{
		// inline function via anonymous function
		func(){
			fmt.Println(idx)
		}()
	}
}
