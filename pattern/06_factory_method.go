package pattern

import "fmt"

const (
	ServType   ComputerType = "server"
	LaptopType ComputerType = "laptop"
)

type ComputerType string

type StoreObject interface {
	GetType() string
	PrintDetails()
}

type Laptopp struct {
	Display  string
	Keyboard string
	Trackpad string
}

type Serv struct {
	CPU    string
	Memory int
}

func NewServ() Serv {
	return Serv{
		CPU:    "Intel",
		Memory: 256,
	}
}

func NewLaptop() Laptopp {
	return Laptopp{
		Display:  "HP",
		Keyboard: "FullSize",
		Trackpad: "Crap",
	}
}

func (n Laptopp) GetType() string {
	return "notebook"
}

func (n Laptopp) PrintDetails() {
	fmt.Printf("Display %s, Keyboard %s, Trackpad %s\n", n.Display, n.Keyboard, n.Trackpad)
}

func (s Serv) GetType() string {
	return "server"
}

func (s Serv) PrintDetails() {
	fmt.Printf("CPU %s, Mem %d\n", s.CPU, s.Memory)
}

func New(typeName ComputerType) StoreObject {
	switch typeName {
	default:
		fmt.Printf("Несуществующий тип %s\n", typeName)
		return nil
	case ServType:
		return NewServ()
	case LaptopType:
		return NewLaptop()
	}
}

// Паттерн фабричный метод применяется для создания объектов разных типов через удобный интерфейс

// Плюсы: простой в использовании интерфейс
// Минусы: есть вероятность превращения в божественный объект, от которого может быть сложно отказаться
