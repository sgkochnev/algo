package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Просматривая локальную сеть одной компании, вы составляете список IP-адресов (разумеется, это IPv7; IPv6 слишком ограничен). Вы хотите выяснить, какие IP-адреса поддерживают SSL (super-secret listening).

// IP поддерживает SSL, если он имеет Area-Broadcast Accessor, или ABA, в любом месте последовательностей суперсети (вне секций, заключенных в квадратные скобки) и соответствующий Byte Allocation Block, или BAB, в любом месте последовательностей гиперсети. ABA - это любая трех символьная последовательность, состоящая из двух одинаковых символов с разными символами между ними, например, xyx или aba.&nbsp; Соответствующий BAB - это те же символы, но в обратных позициях: yxy и bab, соответственно.

// Например:
// 	- aba[bab]xyz поддерживает SSL (aba вне квадратных скобок с соответствующим bab в квадратных скобках).
//  - xyx[xyx]xyx не поддерживает SSL (xyx соответствующий BAB, но нет последовательности yxy соответствующей ABA).
//  - aaa[kek]eke поддерживает SSL (eke в суперсети с соответствующим kek в гиперсети; последовательность aaa не подходит под ABA, потому что внутренний символ должен быть другим).
//  - zazbz[bzb]cdb поддерживает SSL (у zaz нет соответствующего aza, но у zbz есть соответствующий bzb)

// Сколько IP-адресов в вашем списке поддерживают SSL?

func main() {

	f, err := os.OpenFile("007/07.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("file not found")
		log.Fatalln(err)
	}
	defer f.Close()

	fmt.Println(checkFile(f))
}

func checkFile(r io.Reader) int {
	reader := bufio.NewReader(r)
	line, _, err := reader.ReadLine()
	sslSupport := 0
	for ; err != io.EOF; line, _, err = reader.ReadLine() {
		if err != nil {
			fmt.Println("can not read file")
			log.Fatalln(err)
		}
		if checkStr(line) {
			sslSupport++
		}
	}
	return sslSupport
}

func checkStr(str []byte) bool {
	reader := bufio.NewReader(bytes.NewReader(str))
	m := make([][]string, 2)
	m[0] = make([]string, 0)
	m[1] = make([]string, 0)
	for {
		s, err := reader.ReadBytes('[')
		if err == io.EOF {
			m[0] = append(m[0], string(s))
			break
		}
		m[0] = append(m[0], string(s[:len(s)-1]))

		s1, err := reader.ReadBytes(']')
		m[1] = append(m[1], string(s1[:len(s1)-1]))
		if err == io.EOF {
			break
		}
	}

	for _, v := range m[0] {
		for j := 0; j < len(v)-2; j++ {
			xyx := v[j : j+3]
			if aba(xyx) {
				for _, v1 := range m[1] {
					substr := xyx[1:2] + xyx[0:2]
					if strings.Contains(v1, substr) {
						return true
					}
				}
			}
		}
	}
	return false
}

func aba(s string) bool {
	return s[0] == s[2]
}
