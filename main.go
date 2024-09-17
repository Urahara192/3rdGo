package main

import (
	"Third/mathutils"
	"Third/stringutils"
	"fmt"
)

func main() {
	// Чтение числа от пользователя
	var number int
	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&number)

	// Вычисление факториала
	factorial := mathutils.Factorial(number)
	if factorial != -1 {
		fmt.Printf("Факториал %d равен %d\n", number, factorial)
	} else {
		fmt.Println("Факториал не определен для отрицательных чисел.")
	}

	// Пример со строками
	str := "Hello, World!"
	reversedStr := stringutils.Reverse(str)
	fmt.Printf("Перевернутая строка: %s\n", reversedStr)

	// Работа с массивом
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Массив:", array)

	// Создание среза и выполнение операций
	slice := array[1:4] // Создаем срез из массива
	fmt.Println("Срез:", slice)

	// Добавление элемента в срез
	slice = append(slice, 6)
	fmt.Println("Срез после добавления элемента:", slice)

	// Удаление элемента
	slice = append(slice[:1], slice[2:]...) // Удаляем второй элемент
	fmt.Println("Срез после удаления элемента:", slice)

	// Поиск самой длинной строки
	stringSlice := []string{"bratishkin", "evelone", "guacamoly", "mokrivskiy", "shadowkekw"}
	longestString := ""
	for _, s := range stringSlice {
		if len(s) > len(longestString) {
			longestString = s
		}
	}
	fmt.Printf("Самая длинная строка: %s\n", longestString)
}
