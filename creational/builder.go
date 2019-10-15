package creational

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.6
)

type Color string

const (
	BlueColor  Color = "blue"
	RedColor         = "red"
	GreenColor       = "green"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type BuilderInterface interface {
	Color(Color) BuilderInterface
	TopSpeed(Speed) BuilderInterface
	Wheels(Wheels) BuilderInterface
	Build() CarInterface
}

type CarInterface interface {
	Drive() error
	Stop() error
}

// Implementation

type Car struct {
	Color    Color
	Wheels   Wheels
	TopSpeed Speed
}

func (c Car) Drive() error {
	fmt.Print()
	return nil
}

func (c Car) Stop() error {
	return nil
}

type Builder struct {
	car Car
}

func NewBuilder() Builder {
	return Builder{car: Car{}}
}

func (b Builder) Color(color Color) BuilderInterface {
	b.car.Color = color
	return b
}

func (b Builder) Wheels(wheels Wheels) BuilderInterface {
	b.car.Wheels = wheels
	return b
}

func (b Builder) TopSpeed(speed Speed) BuilderInterface {
	b.car.TopSpeed = speed
	return b
}

func (b Builder) Build() CarInterface {
	return b.car
}

// Usage

func builderPattern() {
	builder := NewBuilder()
	sportCar := builder.TopSpeed(300 * KPH).Wheels(SportsWheels).Color(RedColor).Build()
	sportCar.Drive()
}
