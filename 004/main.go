package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// Была введена новая системная политика, которая требует, чтобы все учетные записи использовали кодовую фразу вместо простого пароля.

// Парольная фраза состоит из набора слов (рандомных строчных букв), разделенных пробелами. Для обеспечения безопасности действующая парольная фраза не должна содержать повторяющихся слов.

// Полный список парольных фраз системы вы найдете по кнопке Данные внизу вашей задачи (одна строка из слов - одна парольная фраза). Сколько парольных фраз действительны (т.е. не содержат повторений)?

func main() {
	f, err := os.OpenFile("004/04.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("file not found")
		log.Fatalln(err)
	}
	defer f.Close()

	fmt.Println(runCheck(f))
}

func runCheck(r io.Reader) int {
	reader := bufio.NewReader(r)
	correctPhrase := 0
	line, _, err := reader.ReadLine()
	for err != io.EOF {
		if err != nil {
			fmt.Println("can not read file")
			log.Fatalln(err)
		}
		if checkPhrase(line) {
			correctPhrase++
		}
		line, _, err = reader.ReadLine()
	}
	return correctPhrase
}

func checkPhrase(phrese []byte) bool {
	p := bytes.Split(phrese, []byte(" "))

	for i, word := range p[:len(p)-1] {
		for _, w := range p[i+1:] {
			if bytes.Equal(word, w) {
				return false
			}
		}
	}
	return true
}
