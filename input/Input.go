package input

import (
	"fmt"
	ds "url/dss"
	hash "url/Hashing"
	"strings"
	"bufio"
	"os"
	//"regexp"
)
// Yet to use regexp
func format(s string) string{
	return s
}

func Get_URL(){
	fmt.Println("Enter the URL:")
	ss:=bufio.NewReader(os.Stdin)
	s,_:=ss.ReadString('\n')
	s=strings.TrimSpace(s)
	s=format(s)
	if len(s)>ds.MaxURLLen{
		fmt.Println("URL too long. Try again!!!")
		Get_URL()
	}
	if _,ok:=ds.Ma[s];ok{
		fmt.Println("Shortened URL is:",ds.Ma[s])
		return
	}
	outer:
	for{
	fmt.Println("1. Enter 1 to get a custom URL\n2. Enter 2 to get a random and short URL")
	var b int
	fmt.Scan(&b)
	switch(b){
		case 1: 
			fmt.Println("Enter custom link")
			er,_:=ss.ReadString('\n')
			er="test.com/"+strings.TrimSpace(er)
			ds.Ma[s]=er
			ds.Am[er]=s
			break outer
		case 2: 
			er:=""
			for {
				er=hash.Hash()
				if _,okk:=ds.Am[er];!okk{
					break
				}
			}
			ds.Ma[s]=er
			ds.Am[er]=s
			break outer
		default:
			fmt.Println("Invalid input.Retry again!!!")
	}
}
	fmt.Println("Shortened URL is:",ds.Ma[s])
}
