package main

import (
	"fmt"
	"sync"
	"time"
)

func getUserByName(name string) string {
	_ = name
	time.Sleep(100 * time.Millisecond)
	return "001"
}

func getUserProduct(userID string, respch chan []string, wg *sync.WaitGroup) {
	// when this method is done, it will call wg.Done() to signal that it's done
	defer wg.Done()
	_ = userID
	time.Sleep(400 * time.Millisecond)
	respch <- []string{"p1", "p2", "p3"}
}

func getUserPreference(userID string, respch chan []string, wg *sync.WaitGroup) {
	// when this method is done, it will call wg.Done() to signal that it's done
	defer wg.Done()
	_ = userID
	time.Sleep(300 * time.Millisecond)
	respch <- []string{"a", "b", "c"}
}

func main() {

	start := time.Now()

	userID := getUserByName("gopher")

	productsch := make(chan []string, 1)
	preferencesch := make(chan []string, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go getUserProduct(userID, productsch, &wg)
	go getUserPreference(userID, preferencesch, &wg)

	// wait for both goroutines to finish so when wg is 0
	wg.Wait()
	// close the channels when all the messages are received
	close(productsch)
	close(preferencesch)

	fmt.Println("UserID:", userID)

	products := <-productsch
	fmt.Println("Products:", products)
	preferences := <-preferencesch
	fmt.Println("Preferences:", preferences)

	fmt.Println("Elapsed time:", time.Since(start))

}
