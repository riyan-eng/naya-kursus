package main

import "fmt"

func main() {
	// inisialiasi var
	var data map[string]string
	fmt.Println(data == nil)

	// data["nama"] = "riyan" // akan error karena belum ada nilai awal
	// fmt.Println(data)

	data = map[string]string{}
	data["nama"] = "armada"
	fmt.Println(data)

	// langsung di isi value
	var chicken map[string]int = map[string]int{
		"id":   1,
		"name": 2,
		"age":  10,
	}
	fmt.Println(chicken)

	var tt map[string]string
	fmt.Println(tt)

	var chicken2 = map[string]string{}
	fmt.Println(chicken2)
	chicken2["nama"] = "riri"
	fmt.Println(chicken2)

	for key, val := range chicken {
		fmt.Println(key, val)
	}

	fmt.Println(chicken)
	delete(chicken, "id")
	fmt.Println(chicken)

	var value, ada = chicken["ages"]
	fmt.Println(value)
	fmt.Println(ada)

}
