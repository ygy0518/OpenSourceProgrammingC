package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i) // 줄바꿈등 제거, 파이썬의 strip 함수와 비슷
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	counts := 0
	j := 1
	for j <= n {
		if n%j == 0 {
			counts = counts + 1
		}
		j++
	}

	if counts == 2 {
		fmt.Printf("%d is prime number", n)
	} else {
		fmt.Printf("%d is not prime number", n)
	}
}
