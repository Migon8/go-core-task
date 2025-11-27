package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func concatToString(
	numDecimal, numOctal, numHexadecimal int,
	pi float64,
	name string,
	isActive bool,
	complexNum complex64,
) string {
	return fmt.Sprintf(
		"%d %d %d %f %s %t %v",
		numDecimal, numOctal, numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	)
}

func stringToRunes(s string) []rune {
	return []rune(s)
}

// insertSaltAndHash вставляет соль в середину среза рун и хэширует результат SHA256.
func insertSaltAndHash(runes []rune, salt string) string {
	mid := len(runes) / 2

	saltRunes := []rune(salt)
	salted := make([]rune, 0, len(runes)+len(saltRunes))

	salted = append(salted, runes[:mid]...)
	salted = append(salted, saltRunes...)
	salted = append(salted, runes[mid:]...)

	// Переводим обратно в байты (UTF-8) и считаем SHA256
	data := []byte(string(salted))
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func main() {

	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	fmt.Printf("numDecimal: type=%T, value=%v\n", numDecimal, numDecimal)
	fmt.Printf("numOctal: type=%T, value=%v\n", numOctal, numOctal)
	fmt.Printf("numHexadecimal: type=%T, value=%v\n", numHexadecimal, numHexadecimal)
	fmt.Printf("pi: type=%T, value=%v\n", pi, pi)
	fmt.Printf("name: type=%T, value=%q\n", name, name)
	fmt.Printf("isActive: type=%T, value=%v\n", isActive, isActive)
	fmt.Printf("complexNum: type=%T, value=%v\n\n", complexNum, complexNum)

	combined := concatToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("Combined string:")
	fmt.Println(combined)
	fmt.Println()

	//  Преобразуем строку в срез рун
	runes := stringToRunes(combined)
	fmt.Println("Runes length:", len(runes))

	// Захэшируем этот срез рун SHA256, добавив в середину соль "go-2024"
	const salt = "go-2024"
	hash := insertSaltAndHash(runes, salt)
	fmt.Println("SHA256 hash with salt in the middle:")
	fmt.Println(hash)
}
