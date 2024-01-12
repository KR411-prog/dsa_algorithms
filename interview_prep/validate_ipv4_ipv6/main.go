import (
	"strings"
	"strconv"
	"fmt"
	)

	func isValidIPV4(queryIP string) bool {
	   a := strings.Split(queryIP, ".")
	   if len(a) != 4 {
		   return false
	   }
	   for i:=0;i<4;i++ {
		   if len(a[i]) != 1 && a[i][0] == '0' {
			   return false
		   }
		   temp,err := strconv.Atoi(a[i])
		   if err != nil ||  temp >= 256 || temp < 0 {
			   return false
		   }
	   }

	   return true


	}

	func isValidIPV6(queryIP string) bool {
	  ip := strings.Split(queryIP,":")
	  if len(ip) != 8 {
		  return false
	  }

	  for _,e := range ip {
		  if len(e) <1 || len(e) >4 {
			  fmt.Println("1")
			  return false
		  }
		  for _,j := range e {
			  if !unicode.IsDigit(j) && !unicode.IsLetter(j) {
				  fmt.Println("2")
				  return false
			  }
			  if unicode.IsLetter(j) && (unicode.ToLower(j) < rune('a') || unicode.ToLower(j)>rune('f')) {
					fmt.Println("3")
				  return false
			  }
		  }
	  }
		return true
	}
	func validIPAddress(queryIP string) string {
		if queryIP == "" {
			return "Neither"
		}

		if queryIP[0] == ':' || queryIP[0] == '.' || queryIP[len(queryIP)-1] == ':' || queryIP[len(queryIP)-1] == '.' || strings.Contains(queryIP, "..") || strings.Contains(queryIP, "::") {
			return "Neither"
		}

		if isValidIPV4(queryIP) {
			return "IPv4"
		}
		if isValidIPV6(queryIP) {
				return "IPv6"
			}
		return "Neither"

		}
