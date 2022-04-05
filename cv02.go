package main

import (
	"fmt"
	"math/rand"
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

type queue struct {
	waitingCars []*car
	filledCars  []*car
	payedCars   []*car
}

func newQueue() *queue {
	q := queue{waitingCars: []*car{newCar(gas), newCar(gas), newCar(gas)}}
	return &q
}

type car struct {
	needFuel Fuel
}

func newCar(fuel Fuel) *car {
	c := car{needFuel: fuel}
	return &c
}

type pump struct {
	pumpFuel Fuel
}

func newPump(fuel Fuel) *pump {
	p := pump{pumpFuel: fuel}
	return &p
}

type cashRegister struct {
}

func newCashRegister() *cashRegister {
	cR := cashRegister{}
	return &cR
}

type station struct {
	pumps         []*pump
	cashRegisters []*cashRegister
}

func newStation() *station {
	s := station{}
	/*
		for i := 1; i <= 4; i++ {
			s.pumps = append(s.pumps, newPump(gas))
		}
		for i := 1; i <= 4; i++ {
			s.pumps = append(s.pumps, newPump(diesel))
		}
		for i := 1; i <= 1; i++ {
			s.pumps = append(s.pumps, newPump(lpg))
		}
		for i := 1; i <= 8; i++ {
			s.pumps = append(s.pumps, newPump(electric))
		}
		for i := 1; i <= 2; i++ {
			s.cashRegisters = append(s.cashRegisters, newCashRegister())
		}
	*/
	return &s
}

func usePump(typeOfFuel Fuel, queue *[100]*car, wg *sync.WaitGroup) {
	fmt.Println(typeOfFuel + " ready")
	var wg2 sync.WaitGroup
	for i := 0; i < len(queue); i++ {
		wg2.Add(1)
		go func() {
			time.Sleep(2 * time.Second)
			fmt.Println(queue[i].needFuel + " filled")
			wg2.Done()
		}()

	}
	wg2.Wait()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var queue [100]*car
	gasTypes := []string{"gas", "diesel", "lpg", "electric"}
	for i := 0; i < len(queue); i++ {
		queue[i] = newCar(Fuel(gasTypes[rand.Intn(len(gasTypes)-0)+0]))
	}
	wg.Add(1)
	go usePump("gas", &queue, &wg)
	wg.Add(1)
	go usePump("diesel", &queue, &wg)
	wg.Add(1)
	go usePump("lpg", &queue, &wg)
	wg.Add(1)
	go usePump("electric", &queue, &wg)
	wg.Wait()
	fmt.Println("Done")
}
