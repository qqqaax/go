package main

import (
	"fmt"
	"math/rand"
	"time"
)

//난수 추출된 수의 소수 판정 프로그램v0.3
//소수 : 1과 자신 자신 외에는 나누어 떨어지지 않는 수(0과 1은 제외)
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 // 0과 1 제외, 2 ~ 151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ { // 1과 number일 때 loop 돌지 않음
		if number%i == 0 {
			count = count + 1
		}
	}

	if count == 0 {
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}

}
