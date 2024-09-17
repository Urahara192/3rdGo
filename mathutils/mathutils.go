package mathutils

// Factorial вычисляет факториал числа n.
func Factorial(n int) int {
	if n < 0 {
		return -1 // Возвращаем -1 для отрицательных чисел
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
