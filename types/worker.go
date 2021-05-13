package types

// FuncInterface defines the func used in scheduling.
//	Generally it's better keeping invocation fast but if it costs much more time
//	maybe you should carefully set a suitable timeout during shutdown.
type FuncInterface func(strategyId, parameter string)

// Worker manages data of scheduling for binded strategy
type Worker interface {
	// XXX deal with the error
	// XXX remove the parameters
	Start(strategyId, parameter string) error
	Stop(strategyId, parameter string) error
}