package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var isPrime bool = true
	// bug fix
	if n <= 1 {
		isPrime = false
	} else {
		j := 2
		for j <= int(math.Sqrt(float64(n))) {
			if n%j == 0 {
				isPrime = false
				break // performance up
			}
			fmt.Printf("%d ", j) // Check j loop
			j++
		}
	}
	if isPrime {
		fmt.Printf("%d is prime number", n)
	} else {
		fmt.Printf("%d is not prime number", n)
	}
}
