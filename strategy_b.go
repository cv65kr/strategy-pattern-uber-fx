package main

type StrategyB struct {
}

func NewStrategyB() StrategyB {
	return StrategyB{}
}

func (s StrategyB) handle() string {
	return "CCC"
}

func (s StrategyB) supports(input string) bool {
	return input == "B"
}
