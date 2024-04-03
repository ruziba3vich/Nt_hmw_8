package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fileLen funktsiya fayl nomini oladi va fayldagi baytlar sonini qaytaradi.
	// Agar faylni o'qishda xatolik bo'lsa, xatoni qaytaring.
	// Fayl to'g'ri yopilganligiga ishonch hosil qilish uchun defer dan foydalaning.
	// file oqish uchun os.Open method iborat
	res, err := fileLen("simple.txt")
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("simple.txt fayldagi bytelar soni ", res, " ta")
	}
}

func fileLen(fileName string) (res int, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		/*
		Emojilar qatnashib qolishini hisobga olgan holda alohida hisoblab chiqdim
		aks holda odatiy char 1 byte bo'lar edi va shunchaki bir qator uzunligi `len`
		ni topish yetarli bo'lar edi
		*/
		for _, r := range scanner.Text() {
			res += int(r)
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return res, nil
}
