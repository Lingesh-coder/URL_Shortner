package main
import (
	"fmt"
	get "url/input"
)

func main(){
	outer:
	for{
		var ch rune
		fmt.Println("1. Enter Y/y to get a shortened URL\n2. Enter N/n to exit")
		fmt.Scanf("%c",&ch)
		switch(ch){
		case 'Y','y':
			get.Get_URL()
		case 'N','n':
			fmt.Println("Exiting...")
			break outer
		default:
			fmt.Println("Enter character again")
		}
	}
}