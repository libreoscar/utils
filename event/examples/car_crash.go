package main

import (
	"github.com/libreoscar/utils/event"
	"log"
	"time"
)

type Car struct {
	//-------------------------- events ----------------------------
	Crashed *event.Event

	nPassengers int
}

func NewCar(nPassengers int) *Car {
	return &Car{
		Crashed:     event.NewEvent(),
		nPassengers: nPassengers,
	}
}

func (car *Car) DrunkDriving() {
	log.Printf("Drunk driving with %d passengers\n", car.nPassengers)
	// ...
	car.Crashed.Emit(car.nPassengers / 2)
	// ...
}

func (car *Car) Drown() {
	log.Printf("Car drowned with %d passengers\n", car.nPassengers)
	// ...
	car.Crashed.Emit(car.nPassengers + 1)
	// ...
}

type Reporter struct {
	CarCrashChan chan int
}

func NewReporter() (r *Reporter) {
	r = &Reporter{make(chan int)}
	// r = &Reporter{make(chan int, 10)}  // use a buffered chan when appropriate
	go func() { // a daemon for message processing
		for {
			select {
			case death := <-r.CarCrashChan:
				log.Printf("ChanHandler: Car crashed, %d dead.\n", death)
			}
		}
	}()
	return
}

func main() {
	r := NewReporter()
	car := NewCar(10)

	car.Crashed.Subscribe(func(n int) { // direct call
		log.Printf("FuncHandler: Car crashed, %d dead.\n", n)
	})

	car.Crashed.Subscribe(func(n int) { // multi-thread
		r.CarCrashChan <- n
	})

	car.Crashed.Subscribe(func(n int) { // multi-thread, best effort
		select {
		case r.CarCrashChan <- n:
			return
		default:
			return
		}
	})

	go func() {
		car.DrunkDriving()
		car.Drown()
	}()

	time.Sleep(1000 * time.Millisecond)
}
