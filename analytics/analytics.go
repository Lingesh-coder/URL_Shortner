package analytics

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	ds "url/dss"
)

func Analytics() (string, error) {
	fmt.Println("Enter shortURL to view analytics")
	scan := bufio.NewReader(os.Stdin)
	str, _ := scan.ReadString('\n')
	str = strings.TrimSpace(str)
	if _, ok := ds.Am[str]; !ok {
		return "The short link does not exist", errors.New("404. Not Found")
	}

	return (str + " is viewed " + strconv.Itoa(ds.AnalyticsMap[str]) + " time(s)"), nil
}
