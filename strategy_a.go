package main

type StrategyA struct {
}

func NewStrategyA() StrategyA {
	return StrategyA{}
}

func (s StrategyA) handle() string {
	return "ABC"
}

func (s StrategyA) supports(input string) bool {
	return input == "A"
}
