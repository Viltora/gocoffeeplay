package chapterone

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//Упражнение 1.3. Поэкспериментируйте с измерением разницы времени выполнения
//потенциально неэффективных версий и версии с применением strings.join.
//(В разделе 1.6 демонстрируется часть пакета time, а в разделе 11.4 — как написать
//тест производительности для ее систематической оценки.)

func Task13() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.8fs\n", secs)
	fmt.Println(s)
	//время 0.00001467s

	//решение

	var s1 string
	start1 := time.Now()

	s1 += strings.Join(os.Args[1:], " ")

	secs1 := time.Since(start1).Seconds()
	fmt.Printf("%.8fs\n", secs1)
	fmt.Println(s1)
	//время 0.00000033s
}
