package main

import (
	"fmt"
	"sync"
	"time"
)

func fonksiyon1(wg *sync.WaitGroup) {

	time.Sleep(2 * time.Second)
	fmt.Println("Fonk1 tamamlandı")

	wg.Done() //-1
}

func fonksiyon2(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("Fonk2 tamamlandı")

	wg.Done() //-1
}

func fonksiyon3(wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	fmt.Println("Fonsiyon 3 tanımlandı")

	wg.Done() //-1
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go fonksiyon1(&wg)
	go fonksiyon2(&wg)
	go fonksiyon3(&wg)
	fmt.Println("Merhaba Dünya!")

	wg.Wait() //Toplam değerin 0 olması bekler

	fmt.Println("WaitGroup'lar tamamlandı.")
}
