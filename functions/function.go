package function

import (
	"fmt"
	"os"
	"strings"
)

func CheckFormatCommand() bool {
	arg := os.Args
	if len(arg) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return false
	}
	if arg[2]  != "standard"   && arg[2] != "shadow" && arg[2] != "thinkertoy" {
		fmt.Println("Invalid banner type. Please use: standard, shadow, thinkertoy")
		return false
	}
	for _, val := range os.Args[1] {
		if val < 32 || val > 126 {
			fmt.Println("Invalid string. Try to insert an printable character!")
			return false
		}
	}
	return true
}

func CheckBanner(banner string) bool {
	if  banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
		return true
	}
	return false
}
func ReadArg() (string, string){
	banner := os.Args[1]
	t := os.Args[2]
	text, err := os.ReadFile("./banners/"+banner+".txt")
	if err != nil {
		fmt.Println("The error is => ", err)
	}
	return string(text), t
}

func TraitmentData(bnr string,  arg string) string {
	banner := bnr
	// "../banners/"
	text, err := os.ReadFile("./banners/"+banner+".txt")
	if err != nil {
		fmt.Println("The error is => ", err)
	}
	
	arrData := strings.Split(string(text), "\n")
	arg = strings.ReplaceAll(arg, "\\n", "\n")
	res := ""
	count := 0
	words := strings.Split(arg, "\n")
	for k := 0; k < len(words); k++ {
		if words[k] == "" {
			res += "\n"
			count++
			continue
		}
		
	for i := 0; i < 8; i++ {
		for j := 0; j < len(words[k]); j++ {
			
				start := (int(words[k][j]-' ') * 9) + i + 1
				res +=  arrData[start]
			}
			res +=  "\n"
		}
	}
	if count == len(words) {
		return res[:len(res)-1]
	}
	return res
}