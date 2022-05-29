package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// Санта доставляет подарки в бесконечную двухмерную сетку домов.

// Он начинает с доставки подарка в дом в начальной точке, а затем эльф с Северного полюса вызывает его по радио и сообщает, куда двигаться дальше. Перемещение всегда происходит ровно на один дом на север (^), юг (v), восток (>) или запад (<). После каждого перемещения он доставляет очередной подарок в дом на новом месте.

// Однако эльф, вернувшийся на Северный полюс, немного перебрал гоголь-моголя, поэтому его указания немного сбились, и Санта в итоге посещает некоторые дома более одного раза.; Сколько домов получили хотя бы один подарок?

// Например:
// 	- > доставляет подарки в 2 дома: один в начальной точке, и один на востоке.
// 	- ^>v<; доставляет подарки в 4 дома в квадрате, в том числе дважды в дом в его начальном/конечном местоположении.
// 	- ^v^v^v^v^v^v^v доставляет кучу подарков некоторым очень удачливым детям только в 2 дома.

func main() {
	f, err := os.OpenFile("003/03.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("file not found")
		log.Fatalln(err)
	}
	defer f.Close()

	fmt.Println(run(f))
}

type coordinates struct {
	x int
	y int
}

func run(r io.Reader) int {
	m := make(map[coordinates]int)
	cs := coordinates{}
	m[cs]++
	reader := bufio.NewReader(r)

	b, err := reader.ReadByte()
	for err != io.EOF {
		cs = newCoordinates(cs, b)
		m[cs]++
		b, err = reader.ReadByte()
	}
	return len(m)
}

func newCoordinates(cs coordinates, b byte) coordinates {
	switch b {
	case '^':
		cs.y += 1
	case 'v':
		cs.y -= 1
	case '<':
		cs.x -= 1
	case '>':
		cs.x += 1
	}
	return cs
}
