package main

import "fmt"

func main() {
	ids := []int{12, 53, 22, 41, 90}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using Index
	for _, id := range ids {
		fmt.Printf("ID - %d\n", id)
	}

	// Sum all the ID
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum: %d\n", sum)

	// Range with map
	emails := map[string]string{"felix": "felixroo98@gmail.com", "sarah": "sarah@gmail.com", "wilson": "wilson@gmail.com"}
	fmt.Println(emails)

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// Key only
	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}
