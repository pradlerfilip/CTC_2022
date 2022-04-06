package main

import (
	"fmt"
	"sync"
	"time"
)

type Fuel string

const (
	gas      = "gas"
	diesel   = "diesel"
	lpg      = "lpg"
	electric = "electric"
)

func fillCar(n_of_cars int, fuel Fuel, capacity chan int, car_queue *sync.WaitGroup) {
	var wg2 sync.WaitGroup
	for i := 1; i <= n_of_cars; i++ {
		capacity <- 1
		wg2.Add(1)
		go func() {
			fmt.Println("Fuelling... " + fuel)
			time.Sleep(time.Second * 1)
			fmt.Println("Fuelling: " + fuel + "  completed")
			<-capacity
			wg2.Done()
		}()
	}
	wg2.Wait()
	car_queue.Done()
}

func main() {

	var car_queue sync.WaitGroup
	capacity_gas := make(chan int, 4)
	capacity_diesel := make(chan int, 4)
	capacity_lpg := make(chan int, 4)
	capacity_electric := make(chan int, 2)
	var n_of_gas int = 5
	var n_of_diesel int = 5
	var n_of_lpg int = 5
	var n_of_electric int = 5
	car_queue.Add(1)
	go fillCar(n_of_gas, "gas", capacity_gas, &car_queue)
	car_queue.Add(1)
	go fillCar(n_of_diesel, "diesel", capacity_diesel, &car_queue)
	car_queue.Add(1)
	go fillCar(n_of_lpg, "lpg", capacity_lpg, &car_queue)
	car_queue.Add(1)
	go fillCar(n_of_electric, "electric", capacity_electric, &car_queue)
	car_queue.Wait()
}
