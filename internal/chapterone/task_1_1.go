package chapterone

import (
	"fmt"
	"os"
)

//Упражнение 1.1. Измените программу echo так, чтобы она выводила также
//os.A rgs[0], имя выполняемой команды.

func Task11() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
