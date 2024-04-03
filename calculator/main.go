package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(i int, j int) (int, error) { return i + j, nil }

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) {
	if j == 0 {
		return -1, errors.New("division by zero")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	// Oddiy kalkulyator dasturi bitta xato holatini hal qilmaydi: nolga bo'lish.
	// Int va xatoni qaytarish uchun matematik operatsiyalar uchun funktsiya signature ni o'zgartiring.
	// Div funksiyasida, agar bo'luvchi 0 bo'lsa, xatoliklarni qaytaring.
	// Xato uchun yangi errors.New("division by zero") qaytaring.
	// Boshqa barcha holatlarda, nil qaytaring.
	// Ushbu xatoni tekshirish uchun main funktsiyani o'zgartiring

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
		{"45", "/", "0"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println("Error :", err)
		} else {
			fmt.Println(result)
		}
	}
}
