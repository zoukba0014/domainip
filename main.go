package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "-h" {
			fmt.Printf("-ip 输出ip目标\n")
			fmt.Printf("-domain 输出domain目标\n")
			fmt.Printf("[空] 全部输出\n")
			os.Exit(1)
		} else {
			reader := bufio.NewReader(os.Stdin)
			var target []string
			for {
				line, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				target = append(target, strings.Replace(line, "\n", "", -1))
			}
			if os.Args[1] == "-ip" {
				iplists, _ := RegTarget(target)
				for _, ip := range iplists {
					fmt.Printf("%s\n", ip)
				}
			} else if os.Args[1] == "-domain" {
				_, domainlists := RegTarget(target)
				for _, domain := range domainlists {
					fmt.Printf("%s\n", domain)
				}
			}
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		var target []string
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			target = append(target, strings.Replace(line, "\n", "", -1))
		}
		iplists, domainlists := RegTarget(target)
		fmt.Printf("iplist: \n")
		for _, ip := range iplists {
			fmt.Printf("%s\n", ip)
		}
		fmt.Printf("domainlist: \n")
		for _, domain := range domainlists {
			fmt.Printf("%s\n", domain)
		}
	}
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func RegTarget(targets []string) ([]string, []string) {
	var iplists []string
	var domainlist []string
	ipReg := `^((0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5])\.){3}(0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5])$`
	//chReg := `^[\u4e00-\u9fa5]{3,8}$`
	//count := 0
	for _, a := range targets {
		ipmatch, _ := regexp.MatchString(ipReg, a)
		//chmatch, _ := regexp.MatchString(chReg, a)
		if ipmatch {
			iplists = append(iplists, a)
		} else if !IsChineseChar(a) {
			//fmt.Println(chmatch)
			domainlist = append(domainlist, a)
			//count++
		}
	}
	//fmt.Println(count)
	return iplists, domainlist
}
