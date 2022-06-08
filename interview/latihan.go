package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Println(alternateCase("Hello World"))

	fibonacci(sum(x[:]))
	sum(x[:])

	fmt.Println(reverseWord("i am Great human"))

	fmt.Println(middleAlphabet("a", "c"))

	fmt.Println((3 + 5 + 6 + 9))

	fmt.Println(productArray([]int{12, 20}))

}

//2
func productArray(nums []int) []int {
	var result []int

	prod := 1
	for _, val := range nums {
		prod *= val
	}

	for _, val := range nums {
		result = append(result, int(float64(prod)*math.Pow(float64(val), -1)))
	}

	return result
}

//3
func alternateCase(str string) string {
	var result string
	for _, val := range str {
		alphabet := string(val)
		if alphabet == " " {
			result += " "
			continue
		}

		if alphabet == strings.ToUpper(alphabet) {
			result += strings.ToLower(alphabet)
			continue
		}

		result += strings.ToUpper(alphabet)
	}
	return result
}

//4
func multipleFiveThree(num int) int {
	var result int
	for i := 1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	return result
}

//5
var x = [3]int{15, 1, 3}

func sum(x []int) int {
	hasilsum := 0
	for _, v := range x {
		hasilsum += v
	}

	fmt.Println(hasilsum)
	return hasilsum
}

func fibonacci(n int) int {
	a := 0
	b := 1

	for {
		temp := a
		a = b
		b = temp + a
		fmt.Println("a", a, "b", b)
		if b > n {
			break
		}
	}
	if math.Abs(float64(n-b)) > math.Abs(float64(n-a)) {
		fmt.Println("if", n-a)
	} else if math.Abs(float64(n-a)) > math.Abs(float64(n-b)) {
		fmt.Println("jarak fibonacci adalah: ", math.Abs(float64(n-b)))
	}
	return n
}

//6
func reverseWord(str string) string {
	reverse := func(str string) string {
		var res string
		for _, v := range str {
			res = string(v) + res
		}
		return res
	}

	arrStr := strings.Split(str, " ")
	for idx, val := range arrStr {
		strReverse := reverse(val)
		if val == strings.Title(val) {
			arrStr[idx] = strings.Title(strings.ToLower(strReverse))
			continue
		}
		arrStr[idx] = strReverse
	}
	return strings.Join(arrStr, " ")
}

//7

const alphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func middleAlphabet(firstAlphabet, lastAlphabet string) string {
	firstAlphabet = strings.ToUpper(firstAlphabet)
	lastAlphabet = strings.ToUpper(lastAlphabet)

	// ! without looping
	first := float64(strings.Index(alphabets, firstAlphabet))
	last := float64(strings.Index(alphabets, lastAlphabet))

	max := math.Max(first, last)
	min := math.Min(first, last)
	center := max - min
	alpha := func(pos float64) string {
		return string(alphabets[int(pos)])
	}

	if center == 0 {
		return alpha(max)
	}
	modulo := int(center) % 2
	devide := math.Floor(center / 2)
	if modulo == 0 {
		return alpha(min + devide)
	}

	return alpha(min+devide) + alpha(min+devide+1)
}
