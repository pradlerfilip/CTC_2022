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

func fillCar(n_of_cars int, fuel Fuel, pump_capacity chan int, cash_registers_capacity chan int, car_queue *sync.WaitGroup) {
	var wg2 sync.WaitGroup
	for i := 1; i <= n_of_cars; i++ {
		pump_capacity <- 1
		wg2.Add(1)
		go func() {
			fmt.Println("Fuelling... " + fuel)
			time.Sleep(time.Second * 1)
			fmt.Println("Fuelling: " + fuel + "  completed")
			cash_registers_capacity <- 1
			fmt.Println("Car with: " + fuel + "  is paying...")
			time.Sleep(time.Millisecond * 500)
			fmt.Println("Car with: " + fuel + "  payed")
			wg2.Done()
			<-cash_registers_capacity
			<-pump_capacity
		}()
	}
	wg2.Wait()
	car_queue.Done()
}

func main() {

	var car_queue sync.WaitGroup
	cash_registers_capacity := make(chan int, 2)
	pump_capacity_gas := make(chan int, 4)
	pump_capacity_diesel := make(chan int, 4)
	pump_capacity_lpg := make(chan int, 4)
	pump_capacity_electric := make(chan int, 2)
	var n_of_gas int = 5
	var n_of_diesel int = 5
	var n_of_lpg int = 5
	var n_of_electric int = 5
	car_queue.Add(1)
	go fillCar(n_of_gas, "gas", pump_capacity_gas, cash_registers_capacity, &car_queue)
	car_queue.Add(1)
	go fillCar(n_of_diesel, "diesel", pump_capacity_diesel, cash_registers_capacity, &car_queue)
	car_queue.Add(1)
	go fillCar(n_of_lpg, "lpg", pump_capacity_lpg, cash_registers_capacity, &car_queue)
	car_queue.Add(1)
	go fillCar(n_of_electric, "electric", pump_capacity_electric, cash_registers_capacity, &car_queue)
	car_queue.Wait()
}
