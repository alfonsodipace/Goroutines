package main

import (
	"fmt"
	"time"
)

func getUserByName(name string) string {
	_ = name
	time.Sleep(100 * time.Millisecond)
	return "001"
}

func getUserProducts(userID string) []string {
	_ = userID
	time.Sleep(400 * time.Millisecond)
	return []string{"p1", "p2", "p3"}
}

func getUserPreferences(userID string) []string {
	_ = userID
	time.Sleep(300 * time.Millisecond)
	return []string{"a", "b", "c"}
}

func main() {

	start := time.Now()

	userID := getUserByName("gopher")
	products := getUserProducts(userID)
	preferences := getUserPreferences(userID)

	fmt.Println("UserID:", userID)
	fmt.Println("Products:", products)
	fmt.Println("Preferences:", preferences)

	fmt.Println("Elapsed time:", time.Since(start))

}
