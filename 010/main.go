package main

// Например:
// 	1		->	11			(одна 1)
//	11		->	21			(две 1)
//	21		->	1211		(одна 2 одна 1)
//  1211	->	111221		(одна 1 одна 2 две 1)
// 	111221 	-> 	312211 		(три 1 две 2 одна 1)
// 	312211	->	13112221 	(одна 3 одна 1 две 2 две 1) ...
// Сгенерировать последовательность элементов начиная с 1113222113.
// Вывести длину 40-го сгенерированного элемента.
import "fmt"

func main() {
	s := "1113222113"
	for i := 0; i < 40; i++ {

		if len(s) == 1 {
			s = "1" + s
			continue
		}
		smb := rune(s[0])

		count := 1
		newstr := ""
		for _, v := range s[1:] {
			if v != smb {
				newstr += fmt.Sprintf("%d%s", count, string(smb))
				smb = v
				count = 1
				continue
			}
			count++
		}
		newstr += fmt.Sprintf("%d%s", count, string(smb))

		s = newstr
	}
	fmt.Println(len(s))
}
