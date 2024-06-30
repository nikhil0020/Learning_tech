package store

type Boat struct {
	*Product
	Capacity  int
	Motorized bool
}

// A struct can mix regular and embedded field types, but the embedded fields are an important part of the composition feature

func NewBoat(name string, price float64, capacity int, motorized bool) *Boat {
	return &Boat{
		NewProduct(name, "Watersports", price),
		capacity,
		motorized,
	}
}
