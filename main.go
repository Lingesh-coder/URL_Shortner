package main
import (
	"fmt"
	get "url/input"
	put "url/output"
	ana "url/analytics"
)

func main(){
	outer:
	for{
		var a int
		fmt.Println("1. Enter 1 to get a shortened URL\n2. Enter 2 to retrieve the Original URL\n3. Enter 3 to view analytics\n4. Enter 4 to exit")
		fmt.Scan(&a)
		switch(a){
		case 1:
			get.Get_URL()
		case 2:
			fmt.Println(put.Put_URL())
		case 3:
			fmt.Println(ana.Analytics());
		case 4:
			fmt.Println("Exiting...")
			break outer
		default:
			fmt.Println("Enter character again")
		}
	}
}