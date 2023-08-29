package main

import "go.uber.org/fx"

var StrategyModule = fx.Module("strategy_module",
	fx.Provide(
		fx.Annotate(
			NewContextService,
			fx.ParamTags(`group:"strategy"`),
		),
		asStrategy(NewStrategyA),
		asStrategy(NewStrategyB),
	),
)

func asStrategy(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(StrategyInterface)),
		fx.ResultTags(`group:"strategy"`),
	)
}
