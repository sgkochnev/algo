package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Связисту нужно помочь выяснить, какие строки в его файле являются сломанными или целыми.

// Целая строка — это строка, обладающая всеми следующими свойствами:
// 	- Она содержит пару из любых двух букв, которые встречаются в строке как минимум дважды без перекрытия, например <em>xyxy</em> (<strong>xy</strong>) или aabcdefgaa (aa), но не aaa (aa перекрывается).
// 	- Она содержит по крайней мере одну букву, которая повторяется с ровно одной буквой между повторами, например, xyx, abcdefeghi (efe) или даже aaa.

// Например:
// 	- qjhvhtzxzqqjkmpb целая, потому что в нем есть пара, которая появляется дважды (qj), и буква, которая повторяется ровно с одной буквой между ними (zxz).
// 	- xxyxx целая, потому что в ней есть пара, которая появляется дважды, и буква, которая повторяется с одной буквой между ними, хотя буквы, используемые каждым правилом, перекрываются.
// 	- uurcxstgmygtbstg - сломанный, потому что в нем есть пара (tg), но нет повтора с одной буквой между ними.
// 	- ieodomkazucvgmuy - сломанный, потому что в ней есть повторяющаяся буква с одной между ними (odo), но нет пары, которая появляется дважды.

// Сколько строк являются целыми?

func main() {

	f, err := os.OpenFile("006/06.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("file not found")
		log.Fatalln(err)
	}
	defer f.Close()

	fmt.Println(check(f))
}

func check(r io.Reader) int {
	reader := bufio.NewReader(r)
	n := 0
	line, _, err := reader.ReadLine()

	for ; err != io.EOF; line, _, err = reader.ReadLine() {
		if err != nil {
			fmt.Println("can not read file")
			log.Fatalln(err)
		}
		if checkStr(string(line)) {
			n++
		}
	}
	return n
}

func checkStr(s string) bool {
	pair := s[:2]
	p3 := s[0:3]
	hasPair := false
	hasP3 := isPalindrome3(p3)
	for i := 2; i < len(s)-1; i++ {
		if !hasPair {
			hasPair = strings.Contains(s[i:], pair)
			pair = s[i-1 : i+1]
		}
		if !hasP3 {
			p3 = s[i-1 : i+2]
			hasP3 = isPalindrome3(p3)
		}
	}
	return hasPair && hasP3
}

func isPalindrome3(s string) bool {
	return s[0] == s[2]
}
