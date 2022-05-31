package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Вам передали набор проводов и побитовых логических вентилей.

// Каждый провод имеет идентификатор (несколько строчных букв) и может нести 16-битный сигнал (число от 0 до 65535). Сигнал подается на каждый провод через ворота, другой провод или какое-то определенное значение. Каждый провод может получать сигнал только от одного источника, но может подавать его в несколько мест назначения. Логические вентили не подают сигнал, пока на всех их входах не появится сигнал.

// В прилагаемой инструкции описано, как соединить детали вместе: x AND y -> z означает, что нужно подключить провода x и y к входам AND вентиля, а затем подключить его выход к проводу z.

// Например:
// 	- 123 -> x означает, что сигнал 123 подается на провод x.
// 	- x AND y -> z означает, что побитовое AND из проводов x и y подается на провод z.
// 	- p LSHIFT 2 -> q означает, что значение из провода p сдвигается влево на 2, а затем подается на провод q.
// 	- NOT e -> f означает, что побитовое дополнение значения из провода e передается в провод f.

// Другие возможные вентили включают OR (побитовое ИЛИ) и RSHIFT (сдвиг вправо). Если по какой-то причине вы хотите эмулировать работу логических вентилей, то почти все языки программирования (например, C, JavaScript или Python, Ruby) предоставляют операторы для этих логических вентилей.

// Например, вот простая схема:
// 	123 -> x
// 	456 -> y
// 	x AND y -> d
// 	x OR y -> e
// 	x LSHIFT 2 -> f
// 	y RSHIFT 2 -> g
// 	NOT x -> h
// 	NOT y -> i

// После выполнения это сигналы на проводах:
// 	d: 72
// 	e: 507
// 	f: 492
// 	g: 114
// 	h: 65412
// 	i: 65079
// 	x: 123
// 	y: 456

// Набор инструкций предоставлен вам в Данные.
// Выполните инструкции и найдите сигнал, который получился на проводе а после всех инструкций, после этого подайте этот сигнал на b вместо сигнала прописанного в инструкциях.
// Полученный из а сигнал -> b
// Выполните инструкции заново.

// Какой сигнал в конечном итоге подается на провод a?

type instruction struct {
	in1     string
	command string
	in2     string
	out     string
	flag    bool
}

func main() {

	f, err := os.OpenFile("008/08.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("file not found")
		log.Fatalln(err)
	}
	defer f.Close()

	i := parseFile(f)
	fmt.Println(work(i))
}

func parseFile(r io.Reader) []instruction {
	reader := bufio.NewReader(r)
	line, _, err := reader.ReadLine()
	inst := make([]instruction, 0)
	for ; err != io.EOF; line, _, err = reader.ReadLine() {
		if err != nil {
			fmt.Println("can not read file")
			log.Fatalln(err)
		}
		sl := strings.Split(string(line), " ")
		// fmt.Println(sl)
		i := instruction{}
		switch len(sl) {
		case 3:
			i.in1 = sl[0]
			i.out = sl[2]
		case 4:
			i.command = sl[0]
			i.in1 = sl[1]
			i.out = sl[3]
			inst = append(inst, i)
		case 5:
			i.in1 = sl[0]
			i.command = sl[1]
			i.in2 = sl[2]
			i.out = sl[4]
		}
		inst = append(inst, i)
	}
	return inst
}

func work(instr []instruction) uint16 {
	k := 1
	m := make(map[string]uint16)
	for k > 0 {
		k = 0
		for i, v := range instr {
			if !v.flag {
				switch v.command {
				case "":
					n, err := strconv.Atoi(v.in1)
					if err != nil {
						if val, ok := m[v.in1]; ok {
							m[v.out] = val
							instr[i].flag = true
						} else {
							k++
						}
					} else {
						m[v.out] = uint16(n)
						instr[i].flag = true
					}
				case "NOT":
					if val, ok := m[v.in1]; ok {
						m[v.out] = ^val
						instr[i].flag = true
					} else {
						k++
					}
				case "AND":
					var vvv uint16
					n, err := strconv.Atoi(v.in1)
					if err != nil {
						val1, ok := m[v.in1]
						if !ok {
							k++
							continue
						}
						vvv = val1
					} else {
						vvv = uint16(n)
					}
					val2, ok := m[v.in2]
					if !ok {
						k++
						continue
					}
					m[v.out] = vvv & val2
					instr[i].flag = true
				case "OR":
					val1, ok := m[v.in1]
					if !ok {
						k++
						continue
					}
					val2, ok := m[v.in2]
					if !ok {
						k++
						continue
					}
					m[v.out] = val1 | val2
					instr[i].flag = true
				case "LSHIFT":
					if val, ok := m[v.in1]; ok {
						n, _ := strconv.Atoi(v.in2)
						m[v.out] = val << n
						instr[i].flag = true
					} else {
						k++
					}
				case "RSHIFT":
					if val, ok := m[v.in1]; ok {
						n, _ := strconv.Atoi(v.in2)
						m[v.out] = val >> n
						instr[i].flag = true
					} else {
						k++
					}
				}
			}
		}
	}
	return m["a"]
}
