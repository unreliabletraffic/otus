package hw02_unpack_string

import (
	      "errors"
      	"fmt"
	      "strings"
	      "unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
      	str := ""
      	if s == "" {
	            	return "", nil
	      }
	      for i, res := range s {
	            	if unicode.IsDigit(res) && i == 0 {
		                  	return "", ErrInvalidString
	            	}
	            	if unicode.IsDigit(res) && unicode.IsDigit(rune(s[i+1])) {
			                  return "", ErrInvalidString
	            	}
	            	if unicode.IsDigit(res) && string(s[i]) == "0" {
		                  	str = str[:len(str)-1]
	            	}
	            	if unicode.IsDigit(res) && s[i-2] == '\\' {
		                  	str += strings.Repeat(string(s[i-2])+string(s[i-1]), int(res-'0')-1)
		                  	continue
            		} else if unicode.IsDigit(res) {
	                  		str += strings.Repeat(string(s[i-1]), int(res-'0')-1)
            		}
            		str += string(s[i])
     	}
    	return str, nil
}
