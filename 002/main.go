package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// Вы идете по коридору и вам кричит светящаяся гуманоидная фигура. "Помогите нам исправить повреждение в этой электронной таблице на двери - если мы не исправим, то дверь заблокируется навсегда!"

// Электронная таблица состоит из строк случайных чисел (доступных по кнопке Данные). Для того чтобы убедиться, что процесс восстановления повреждения идет правильно, им нужно, чтобы вы вычислили контрольную сумму электронной таблицы. Для каждой строки определите разницу между наибольшим значением и наименьшим значением. Контрольная сумма представляет собой сумму всех этих различий. Полученная контрольная сумма будет ответом к этой задаче.

func main() {
	f, err := os.OpenFile("002/02.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("file not found")
		log.Fatalln(err)
	}
	defer f.Close()

	fmt.Println(run(f))
}

func run(rd io.Reader) int {
	r := bufio.NewReader(rd)
	var b []byte
	var err error
	sum := 0
	for ; err != io.EOF; b, _, err = r.ReadLine() {
		if err != nil {
			fmt.Println("can not read file")
			log.Fatalln(err)
		}
		sum += calc(b)
	}
	return sum
}

func calc(s []byte) int {
	numsBytes := split(s, "\t")
	max := scanNum(numsBytes[0])
	min := max
	for _, n := range numsBytes[1:] {
		num := scanNum(n)
		if max < num {
			max = num
			continue
		}
		if min > num {
			min = num
			continue
		}
	}
	return max - min
}

func split(s []byte, sep string) [][]byte {
	return bytes.Split(s, []byte(sep))
}

func scanNum(b []byte) int {
	r := bytes.NewReader(b)
	var n int
	fmt.Fscanf(r, "%d", &n)
	return n
}
