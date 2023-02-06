package main

import "fmt"

func romanToInt(s string) int {
	kamus := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	length := len(s)

	if length == 0 {
		return 0
	}

	jumlah := 0
	i := length - 1
	for i >= 0 {
		if i > 0 && kamus[s[i]] > kamus[s[i-1]] {
			jumlah += kamus[s[i]] - kamus[s[i-1]]
			i -= 2
			continue
		} else {
			jumlah += kamus[s[i]]
		}
		i--
	}
	return jumlah
}

func main() {
	// romanToInt("MCMXCIV")
	fmt.Println("1994")
	fmt.Println(romanToInt("LVIII"))
}
