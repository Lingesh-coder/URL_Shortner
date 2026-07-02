package main
import (
	"fmt"
	get "url/input"
	put "url/output"
)

func main(){
	outer:
	for{
		var a int
		fmt.Println("1. Enter 1 to get a shortened URL\n2. Enter 2 to retrieve the Original URL\n3. Enter 3 to exit")
		fmt.Scan(&a)
		switch(a){
		case 1:
			get.Get_URL()
		case 2:
			fmt.Println(put.Put_URL())
		case 3:
			fmt.Println("Exiting...")
			break outer
		default:
			fmt.Println("Enter character again")
		}
	}
}