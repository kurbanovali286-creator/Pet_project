package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func pl1(ch chan string, wg *sync.WaitGroup) {
	for {
		val := <-ch
		superPunch := rand.IntN(100)
		if superPunch < 20 {
			ch <- "win pl1"
			return
		} else {

			switch val {

			case "ping":
				{
					fmt.Println(val)
					ch <- "pong"
				}
			case "stop":
				{
					wg.Done()

					return
				}

			}
		}

	}
}

func pl2(ch chan string, wg *sync.WaitGroup) {

	for {
		val, _ := <-ch
		superPunch := rand.IntN(100)
		if superPunch < 20 {
			ch <- "win pl2"
			return
		} else {
			switch val {

			case "pong":
				{
					fmt.Println(val)
					ch <- "ping"
				}
			case "stop":
				{
					wg.Done()

					return
				}

			}
		}
	}
}

func main() {

	var point1, point2 int

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Игра началась.")
	ch := make(chan string)

	go pl1(ch, &wg)
	go pl2(ch, &wg)

	var begin int = rand.IntN(2) + 1
	if begin == 1 {
		ch <- "pong"
		fmt.Println("Goes player 1")
	} else {
		fmt.Println("Goes player 2")
		ch <- "ping"
	}

	for {
		val := <-ch
		if val == "win pl1" {
			point1++
		}
		if val == "win pl2" {
			point2++
		}
		if point1 == 14 {
			fmt.Println("Победил Pl1.")
			close(ch)
		}
		if point2 == 14 {
			fmt.Println("Победил Pl2")
			close(ch)

		}
	}

	wg.Wait()
}
