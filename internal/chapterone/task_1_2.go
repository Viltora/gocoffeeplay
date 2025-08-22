package chapterone

import (
	"fmt"
	"os"
)

//Упражнение 1.2. Измените программу echo так, чтобы она выводила индекс и
//значение каждого аргумента по одному аргументу в строке.

func Task12() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
