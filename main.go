package main

import "fmt"

func main() {
	age := 17 // Замените на возраст пользователя

	if isAdult(age) {
		fmt.Print("Пользователь совершеннолетний")
	} else {
		fmt.Print("Пользователь не является совершеннолетним")
	}
}

func isAdult(a int) bool {
	if a >= 18 {
		return true
	} else {
		return false
	}
}
