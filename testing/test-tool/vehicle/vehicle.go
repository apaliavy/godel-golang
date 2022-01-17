package vehicle

type Engine struct {
	name string
}

type Wheel struct {
	color string
}

type Vehicle struct {
	engine *Engine
	wheels []Wheel
}

func NewVehicle(engine *Engine, wheels []Wheel) (*Vehicle, error) {
	//if engine == nil {
	//	return nil, fmt.Errorf("failed to create a vehicle - the engine is missing")
	//}

	return &Vehicle{
		engine: engine,
		wheels: wheels,
	}, nil
}
