package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

// Вы находитесь в городе, где все дома и улицы образуют идеальную сетку, чтобы пройти из пункта А в пункт B вы можете ходить только по улицам и перекресткам, нет никаких переулков, чтобы сократить путь по диагонали или около того.

// У вас есть документ (доступный по нажатию кнопки Данные) с инструкциями с помощью которых вы можете прийти в дом, где проходит вечеринка.

// В документе указано, что вы должны начать с заданных координат (где вы находитесь сейчас) и смотреть на север. Затем следуйте указанной последовательности: поверните налево (L) или направо (R) на 90 градусов, затем пройдите вперед заданное количество блоков, закончив на новом перекрестке.

// К сожалению, инструкции написаны не оптимально и если вы им будете следовать — то не успеете на вечеринку. Учитывая, что ходить можно только по сетке улиц города, как далеко находится кратчайший путь до вечеринки?

// Например:
//  - Следуя за R2, L3 оставляет вас в 2 кварталах на восток и 3 кварталах на север или в 5 кварталах от вас.
//  - R2, R2, R2 оставляет вас в 2 кварталах к югу от вашей начальной позиции, которая находится в 2 кварталах от вас.
//  - R5, L5, R5, R3 оставляет вас в 12 кварталах от вас.

func main() {
	res := 0
	r := calc()
	b := readFile("01.txt")
	steps := bytesToSteps(b, ", ")
	for _, s := range steps {
		res = r(s)
	}
	fmt.Println(res)
}

func readFile(filename string) []byte {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("file not found")
		log.Fatalln(err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("can not read file")
		log.Fatalln(err)
	}
	return b
}

func bytesToSteps(b []byte, sep string) [][]byte {
	return bytes.Split(b, []byte(sep))
}

func scanStep(b []byte) (turn byte, n int) {
	r := bytes.NewReader(b)
	_, err := fmt.Fscanf(r, "%1c%d", &turn, &n)
	if err != nil {
		fmt.Println("can not read step")
		log.Fatalln(err)
	}
	return
}

func calc() func(step []byte) int {
	x, y := 0, 0
	deg := 0
	return func(step []byte) int {
		t, n := scanStep(step)
		switch t {
		case 'L':
			deg = calcDeg(deg, -90)
			x, y = calcСoordinates(deg, x, y, n)
		case 'R':
			deg = calcDeg(deg, 90)
			x, y = calcСoordinates(deg, x, y, n)
		}
		return int(math.Abs(float64(x))) + int(math.Abs(float64(y)))
	}
}

func calcDeg(deg, turn int) int {
	r := 360
	if deg+turn < 0 {
		return r + turn
	}
	return (deg + turn) % r
}

func calcСoordinates(deg, x, y, n int) (int, int) {
	switch deg {
	case 0:
		return x, y + n
	case 90:
		return x + n, y
	case 180:
		return x, y - n
	case 270:
		return x - n, y
	}
	return x, y
}
