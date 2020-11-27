package main

import (
	"github.com/eranet/rhombus/rhomgo"
)

type Position struct {
	Value float64
}

func main() {
	c := rhomgo.LocalJSONConnection()
	defer c.Close()

	robot := "mopla"
	joints := []string{ "left_front_wheel_hinge","right_front_wheel_hinge","left_rear_wheel_hinge","right_rear_wheel_hinge"}

	rate := rhomgo.NewRate(200)
	for x := 0; x< 10;x++ {
		for i := 0.0; i < 10; i += 0.01 {

			for _, name := range joints {
				err := c.Publish(robot+"/"+name+"/command", Position{Value: i})
				if err != nil {
					println("error pub:", err)
				}
			}
			rate.Sleep()
		}
		for i := 1.0; i > -1; i -= 0.01 {

			for _, name := range joints {
				err := c.Publish(robot+"/"+name+"/command", Position{Value: i})
				if err != nil {
					println("error pub:", err)
				}
			}
			rate.Sleep()
		}
	}
/*	for i := -1.0; i < 1; i += 0.001 {

		for _, name := range joints {
			err := c.Publish(robot+"/"+name+"/command", Position{Value: i})
			if err != nil {
				println("error pub:", err)
			}
		}
		rate.Sleep()
	}*/
}
