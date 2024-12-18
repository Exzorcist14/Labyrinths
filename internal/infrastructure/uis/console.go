package uis

import (
	"bufio"
	"fmt"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

const (
	DimensionsInputMessage       = "Введите ширину и высоту базового лабиринта:"
	ErrorDimensionsInputMessage  = "Пожалуйста, введите корректные ширину и высоту базового лабиринта:"
	NoteMessage                  = "Примечание: начало координат лежит в левом верхнем углу, координаты начинаются с нуля"
	StartInputMessage            = "Введите координаты начальной точки:"
	EndInputMessage              = "Введите координаты конечной точки:"
	ErrorCoordinatesInputMessage = "Пожалуйста, введите корректные координаты:"
)

type reader interface {
	Read(p []byte) (n int, err error)
}

type writer interface {
	Write(p []byte) (n int, err error)
	Flush() error
}

// console - консольная реализация пользовательского интерфейса.
type console struct {
	reader   reader
	writer   writer
	renderer renderer
}

// newConsole возвращает указатель на инициализированный console.
func newConsole(r renderer) *console {
	return &console{
		reader:   bufio.NewReader(os.Stdin),
		writer:   bufio.NewWriter(os.Stdout),
		renderer: r,
	}
}

// AskMazeDimensions cпрашивает ширину и высоту.
func (c *console) AskMazeDimensions() (height, width int) {
	areValid := func(data ...any) bool {
		if len(data) != 2 {
			return false
		}

		number1, ok1 := data[0].(*int)
		number2, ok2 := data[1].(*int)

		if !ok1 || !ok2 || *number1 <= 0 || *number2 <= 0 {
			return false
		}

		return true
	}

	AskCorrectData(
		c.printf,
		c.read,
		areValid,
		"%s\n",
		DimensionsInputMessage,
		ErrorDimensionsInputMessage,
		&width, &height,
	)

	return height, width
}

// AskCoordinates cпрашивает координаты start и end.
func (c *console) AskCoordinates(height, width int) (start, end cells.Coordinates) {
	var x, y int

	areValid := func(data ...any) bool {
		if len(data) != 2 {
			return false
		}

		number1, ok1 := data[0].(*int)
		number2, ok2 := data[1].(*int)

		if !ok1 || !ok2 || *number1 < 0 || *number1 >= width || *number2 < 0 || *number2 >= height {
			return false
		}

		return true
	}

	c.printf("\n%s\n", NoteMessage)

	AskCorrectData(
		c.printf,
		c.read,
		areValid,
		"%s\n",
		StartInputMessage,
		ErrorCoordinatesInputMessage,
		&x, &y,
	)

	start = cells.Coordinates{
		X: x,
		Y: y,
	}

	AskCorrectData(
		c.printf,
		c.read,
		areValid,
		"\n%s\n",
		EndInputMessage,
		ErrorCoordinatesInputMessage,
		&x, &y,
	)

	end = cells.Coordinates{
		X: x,
		Y: y,
	}

	return start, end
}

// DisplayMaze отображает лабиринт.
func (c *console) DisplayMaze(mz maze.Maze) {
	c.printf("\n%s\n", c.renderer.Render(mz))
}

// DisplayMazeWithPath отображает лабиринт и путь на нём.
func (c *console) DisplayMazeWithPath(mz maze.Maze, path []cells.Coordinates) {
	c.printf("\n%s\n", c.renderer.RenderPath(mz, path))
}

// AskCorrectData спрашивает данные до тех пор, пока они не будут корректными, читая их в data...;
// данные, которые нужно спросить, должны передаваться по указателю.
func AskCorrectData(
	printf func(format string, a ...any),
	read func(data ...any) error,
	areValid func(data ...any) bool,
	messageFormat, message, errorMessage string,
	data ...any,
) {
	printf(messageFormat, message)

	for {
		errRead := read(data...)

		if errRead != nil || !areValid(data...) {
			printf("%s", errorMessage)
		} else {
			break
		}
	}
}

// read читает данные в data.
func (c *console) read(data ...any) error {
	_, err := fmt.Fscan(c.reader, data...)

	if err != nil {
		return fmt.Errorf("can`t scan data: %w", err)
	}

	_, err = fmt.Fscanln(c.reader)
	if err != nil {
		return fmt.Errorf("can`t scan data: %w", err)
	}

	return nil
}

// printf выводит данные согласно формату.
func (c *console) printf(format string, a ...any) {
	fmt.Fprintf(c.writer, format, a...)
	c.writer.Flush()
}
