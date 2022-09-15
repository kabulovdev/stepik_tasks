package main

import "fmt"

type Car interface {
	move()
	stop()
	enginon()
	enginOFF()
}

func movecar(p Car, ass int) {
	p.move()
}
func stopcar(p Car) {
	p.stop()
}
func enginON(p Car) {
	p.enginon()
}
func enginoff(p Car) {
	p.enginOFF()
}

type Gentra struct {
	model string
	engin string
	oil   int
	km    int
}

func (d *Gentra) move(run int) {
	if d.engin == "ON" {
		fmt.Println("car is start move")
	} else {
		fmt.Println("engin off pls turn on engin")
	}
	if run == 0 {
		fmt.Println("are you kidding me")
	} else {
		for i := 0; i < run; i++ {
			d.oil = d.oil - 1
			d.km = d.km + 1
		}
	}
}
func (d *Gentra) stop() {
	if d.engin == "OFF" {
		fmt.Println("STOP")
	} else {
		fmt.Println("engin off pls turn off engin")
	}
}
func (d *Gentra) enginOFF() {
	if d.engin == "ON" {
		d.engin = "OFF"
		fmt.Println("car engin OFF")
	} else {
		fmt.Println("Car is stop now")
	}
}
func (d *Gentra) enginon() {
	if d.engin == "OFF" {
		d.engin = "ON"
		fmt.Println("car engin ON")
	} else {
		fmt.Println("Car is engin ON now")
	}
}
func (d *Gentra) showall() {
	fmt.Println(d.model)
	fmt.Println(d.engin)
	fmt.Println(d.km)
	fmt.Println(d.oil)
}

func main() {
	car := &Gentra{
		model: "Gentra 1.5",
		engin: "OFF",
		km:    0,
		oil:   100,
	}

	for {
		var a int
		fmt.Println("1.turn on engin.\n2.turn off engin.\n3.move.\n4.stop.\n5.Showall\n")
		fmt.Println("if you want exit enter[6]")
		fmt.Scanln(&a)
		if a == 6 {
			break
		} else if a == 1 {
			car.enginon()
		} else if a == 2 {
			car.enginOFF()
		} else if a == 3 {
			var kmm int
			fmt.Println("qancha km yurish kerak")
			fmt.Scanln(&kmm)
			car.move(kmm)
		} else if a == 4 {
			car.stop()
		} else if a == 5 {
			car.showall()
		}
	}

}
