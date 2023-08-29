package main

type StrategyInterface interface {
	handle() string
	supports(input string) bool
}

type Context struct {
	strategies []StrategyInterface
}


func NewContextService(strategies []StrategyInterface) Context {
	return Context{strategies: strategies}
}

func (c Context) resolve(input string) *string {
	for _, strategy := range c.strategies {
		if strategy.supports(input) {
			val := strategy.handle()
			return &val
		}
	}

	return nil
}
