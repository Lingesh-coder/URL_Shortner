package output

import (
	ds "url/dss"
	"os"
	"strings"
	"bufio"
	"fmt"
	er "url/errors"
)

func Put_URL() string {
	fmt.Println("Enter shortened URL:")
	ss:=bufio.NewReader(os.Stdin)
	s,_:=ss.ReadString('\n')
	s=strings.TrimSpace(s)
	if !ds.ShortURLRegex.MatchString(s){
		fmt.Println("Use test.com/ and then the 6 character URL")
		return Put_URL()
	}
	xd:=er.LengthShortURL(s)
	if xd!=nil{
		fmt.Println(xd)
		return Put_URL()
	}
	xd=er.SpacePresent(s)
	if xd!=nil{
		fmt.Println(xd)
		return Put_URL()
	}
	if y,ok:=ds.Am[s];ok{
		ds.AnalyticsMap[s]++
		return "Original URL is: "+y
	}
	return "NO URL found"
}