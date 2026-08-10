package main

import (
	"fmt"
	get "url/input"
	put "url/output"
	ana "url/analytics"
	er "url/errors"
)

func main(){
	outer:
	for {
		fmt.Println("1. Enter 1 to get a shortened URL\n2. Enter 2 to retrieve the Original URL\n3. Enter 3 to view analytics\n4. Enter 4 to exit")
		a,err:=er.TextToInteger()
		if(err!=nil){
			fmt.Println(err)
			continue
		}
		switch(a){
		case 1:
			get.Get_URL()
		case 2:
			fmt.Println(put.Put_URL())
		case 3:
			xd,dx:=ana.Analytics()
			if dx!=nil{
				fmt.Println(dx)
			} else{
				fmt.Println(xd);
			}
		case 4:
			fmt.Println("Exiting...")
			break outer
		default:
			fmt.Println("Enter 1 to 4. Try again!!")
		}
	}
}
