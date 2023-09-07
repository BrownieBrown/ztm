package challenges

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	array := make([]int, n+1)
	array[0] = 0
	array[1] = 1

	for i := 2; i <= n; i++ {
		array[i] = array[i-1] + array[i-2]
	}

	return array[n]
}

//func fibonacci(n int) int {
//	if n <= 1 {
//		return n
//	}
//	return fibonacci(n-1) + fibonacci(n-2)
//}
