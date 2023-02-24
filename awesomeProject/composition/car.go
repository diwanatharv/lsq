package composition

// struct inside another structure
type engine struct {
	hp int
}

type wheel struct {
	wheeldimension int
}

type car struct {
	CarName string
	Engine  engine
	Wheel   wheel
}

func Newcar(carName string, hp int, wd int) {
	return {
		CarName:carName,
		engine:engine{
          hp
		},

	}
}
