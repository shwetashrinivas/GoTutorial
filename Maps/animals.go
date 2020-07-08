package main

import "fmt"

func main() {
	m := map[string]string{
		"dog": "bark",
	}

	changeMap(m)

	fmt.Println(m)
	/*output:
	map[dog:bark cat:purr]
	*/
}

func changeMap(m map[string]string) {
	m["cat"] = "purr"
}
