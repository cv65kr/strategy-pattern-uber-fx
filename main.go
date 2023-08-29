package main

import (
	"fmt"

	"go.uber.org/fx"
)

func main() {
	options := []fx.Option{
		StrategyModule,
		fx.Invoke(
			func(c Context) {
				result := c.resolve("A")
				if result != nil {
					fmt.Printf("Found strategy, output = " + *result)
				} else {
					fmt.Printf("Strategy not found")
				}
			},
		),
	}

	fx.New(options...).Run()
}
