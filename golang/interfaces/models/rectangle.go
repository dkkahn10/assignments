package models

// Rectangle is a blueprint of a rectangle
type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Length + 2*r.Width
}
