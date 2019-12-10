package main

import (

	"fmt"
	"log"

	"periph.io/x/periph/conn/i2c/i2creg"
    "periph.io/x/periph/conn/physic"
    "periph.io/x/periph/devices/bmxx80"
    "periph.io/x/periph/host"
)

func main(){
	if _,err := host.Init(); err !=nil {
		log.Fatal(err)
		}
		bus,err := i2creg.Open("")
		if err !=nil {
			log.Fatal()
		}
		defer bus.Close()

		dev,err := bmxx80.NewI2C(bus,0x76,&bmxx80.DefaultOpts)

		if err !=nil{
			log.Fatal(err)
		}
		defer dev.Halt()

		var env physic.Env
		if err = dev.Sense($env); err !=nil {

			log.Fatal(err)
		}
		fmt.Printf("%8s %10s %9s\n", env.Temperature, env.Pressure, env.Humidity)
}
