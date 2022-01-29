/*
*       Create on : 2022/1/29 10:56
*       @Author : czx
 */
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(os.Args)

	echo2()
	echo3()
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[0:1] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[0:1], " "))

}
