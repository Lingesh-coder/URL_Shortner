package errors
import (
	"errors"
	ds "url/dss"
)
func Format(s string) error {
	if ds.DomainCheck.MatchString(s){
		return nil
	} else{
		return errors.New("Format error")
	}
}