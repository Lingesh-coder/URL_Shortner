package analytics

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	ds "url/dss"
	"errors"
)

func Analytics() (string,error) {
	fmt.Println("Enter shortURL to view analytics")
	scan:=bufio.NewReader(os.Stdin)
	s,_:=scan.ReadString('\n')
	s=strings.TrimSpace(s)
	if _,ok:=ds.Am[s];!ok{
		return "The short link does not exist",errors.New("404. Not Found")
	}
	return (s + " is viewed "+string(ds.AnalyticsMap[s]+48) + " time(s)"),nil
}
