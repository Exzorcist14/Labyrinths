package cells

// Cell содержит тип клетки и слайс координат клеток, к которым есть переход.
type Cell struct {
	Type        Type
	Transitions []Coordinates
}

// Type описывает тип клетки, храня информацию о её весе.
type Type int

// Coordinates хранит X и Y координаты клетки.
type Coordinates struct {
	X int
	Y int
}

// Константы типа клетки.
const (
	Wall        Type = iota // Стена.
	LightedPass             // Освещённый проход.
	Pass                    // Обычный проход.
)

// Types хранит информацию о доступных типах.
var Types = []Type{Wall, LightedPass, Pass}
