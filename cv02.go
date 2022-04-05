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

//func newQueue() *queue {
//q := queue{waitingCars: []*car{newCar(gas), newCar(gas), newCar(gas)}}
//return &q
//}

type car struct {
	id       int
	needFuel Fuel
}

func newCar(fuel Fuel, id int) *car {
	c := car{needFuel: fuel, id: id}
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
	id   int
	used bool
}

func newCashRegister(id int) *cashRegister {
	cR := cashRegister{used: false, id: id}
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

func usePump(typeOfFuel Fuel, queue *[100]*car, cash_registers *[2]*cashRegister, wg *sync.WaitGroup, channel chan int) {
	defer wg.Done()
	fmt.Println(typeOfFuel + " ready")
	var wg2 sync.WaitGroup
	var wg3 sync.WaitGroup
	for i := 0; i < len(queue); i++ {
		if queue[i].needFuel == typeOfFuel {
			wg2.Add(1)
			go func() {
				time.Sleep(1 * time.Second)
				fmt.Print(queue[i].id)
				fmt.Println("of type: " + queue[i].needFuel + " is filled")
				wg3.Add(1)
				for j := 0; j < len(cash_registers); j++ {
					if cash_registers[j].used == false {
						cash_registers[j].used = true
						time.Sleep(2 * time.Second)
						fmt.Println("of type: " + queue[i].needFuel + " payed")
						cash_registers[j].used = false
						wg3.Done()
					}
					wg3.Wait()
				}
				wg2.Done()
			}()
			wg2.Wait()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var queue [100]*car
	var cash_registers [2]*cashRegister
	channel := make(chan int, 2)
	gasTypes := []string{"gas", "diesel", "lpg", "electric"}
	for i := 0; i < len(queue); i++ {
		queue[i] = newCar(Fuel(gasTypes[rand.Intn(len(gasTypes)-0)+0]), i)
	}
	for i := 0; i < len(cash_registers); i++ {
		cash_registers[i] = newCashRegister(i)
	}
	wg.Add(1)
	go usePump("gas", &queue, &cash_registers, &wg, channel)
	wg.Add(1)
	go usePump("diesel", &queue, &cash_registers, &wg, channel)
	wg.Add(1)
	go usePump("lpg", &queue, &cash_registers, &wg, channel)
	wg.Add(1)
	go usePump("electric", &queue, &cash_registers, &wg, channel)
	wg.Wait()
	fmt.Println("Done")
}
