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
	scan := bufio.NewReader(os.Stdin)
	str,_ := scan.ReadString('\n')
	str = strings.TrimSpace(str)
	if _,ok := ds.Am[str] ; !ok{
		return "The short link does not exist",errors.New("404. Not Found")
	}
	
	return (str + " is viewed "+string(ds.AnalyticsMap[str]+48) + " time(s)"),nil
}
