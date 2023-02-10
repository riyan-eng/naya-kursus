package main

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	la := []int{2, 3, 5}
	for _, val := range la {
		for n%val == 0 {
			n /= val
		}
	}
	return n == 1
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	var year int
	for year = 0; p0 < p; year++ {
		p0 += int(float64(p0)*(percent/100)) + aug
	}
	return year
}

func main() {

	// fmt.Println(isUgly(6))
	// fmt.Println(isUgly(1))
	// fmt.Println(isUgly(14))
	// fmt.Println(isUgly(20))

	NbYear(1500000, 2.5, 10000, 2000000)
	NbYear(1500, 5, 100, 5000)
}
