package generator

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

var (
	dx = []int{-1, 0, 0, 1} // Слайс сдвигов по x, ведущий к смежным по стороне координатам.
	dy = []int{0, -1, 1, 0} // Слайс сдвигов по y, ведущий к смежным по стороне координатам.
)

// IsInside возвращает true, если координаты находятся в пределах лабиринта, иначе false.
func IsInside(coords cells.Coordinates, height, width int) bool {
	return coords.X >= 0 && coords.X < width && coords.Y >= 0 && coords.Y < height
}

// GetRandomCoords возвращает случайные координаты лабиринта.
func GetRandomCoords(height, width int) (cells.Coordinates, error) {
	x, err := GetRandomInt(width)
	if err != nil {
		return cells.Coordinates{}, fmt.Errorf("can`t generate random x coorditane: %w", err)
	}

	y, err := GetRandomInt(height)
	if err != nil {
		return cells.Coordinates{}, fmt.Errorf("can`t generate random y coorditane: %w", err)
	}

	return cells.Coordinates{X: x, Y: y}, nil
}

// GetRandomCoordsFrom возвращает случайные координаты из mp.
func GetRandomCoordsFrom[T any](mp map[cells.Coordinates]T) (cells.Coordinates, error) {
	number, err := GetRandomInt(len(mp))
	if err != nil {
		return cells.Coordinates{},
			fmt.Errorf("can`t generate random number of available coordinates: %w", err)
	}

	cs := cells.Coordinates{}
	i := 0

	for coords := range mp {
		if i == number {
			cs = coords
			break
		}

		i++
	}

	return cs, nil
}

// GetRandomAdjacentCoords возвращает случайные координаты, смежные по стороне с coords.
func GetRandomAdjacentCoords(coords cells.Coordinates, height, width int) (cells.Coordinates, error) {
	nextCoords := coords

	for nextCoords == coords || !IsInside(nextCoords, height, width) {
		offsetNumber, err := GetRandomInt(len(dx))
		if err != nil {
			return cells.Coordinates{}, fmt.Errorf("can`t generate random number of offset: %w", err)
		}

		nextCoords = cells.Coordinates{X: coords.X + dx[offsetNumber], Y: coords.Y + dy[offsetNumber]}
	}

	return nextCoords, nil
}

// GetRandomSignificantType возвращает случайный значимый тип (тип, значение которого больше 0).
func GetRandomSignificantType() (cells.Type, error) {
	var significantTypes []cells.Type

	for _, t := range cells.Types {
		if t > 0 {
			significantTypes = append(significantTypes, t)
		}
	}

	number, err := GetRandomInt(len(significantTypes))
	if err != nil {
		return cells.Pass, fmt.Errorf("can`t generate random number of significant types: %w", err)
	}

	return significantTypes[number], nil
}

// GetRandomInt возвращает случайное число из полуинтервала [0, limit).
func GetRandomInt(limit int) (int, error) {
	result, err := rand.Int(rand.Reader, big.NewInt(int64(limit)))
	if err != nil {
		return 0, fmt.Errorf("can`t generate random int: %w", err)
	}

	return int(result.Int64()), nil
}
