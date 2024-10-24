package ui

import (
	"bufio"
	"fmt"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/renderer"
)

const (
	DIMENSIONS_INPUT_MESSAGE        = "Введите ширину и высоту базового лабиринта:"
	ERROR_DIMENSIONS_INPUT_MESSAGE  = "Пожалуйста, введите корректные ширину и высоту базового лабиринта:"
	NOTE_MESSAGE                    = "Примечание: начало координат лежит в левом верхнем углу, координаты начинаются с нуля"
	START_INPUT_MESSAGE             = "Введите координаты начальной точки:"
	END_INPUT_MESSAGE               = "Введите координаты конечной точки:"
	ERROR_COORDINATES_INPUT_MESSAGE = "Пожалуйста, введите корректные координаты:"
)

// console - консольная реализация пользовательского интерфейса.
type console struct {
	reader   bufio.Reader
	writer   bufio.Writer
	renderer renderer.Renderer
}

// newConsole возвращает указатель на инициализированную структуру консоли.
func newConsole(rendererType string, palette renderer.Palette) *console {
	return &console{
		reader:   *bufio.NewReader(os.Stdin),
		writer:   *bufio.NewWriter(os.Stdout),
		renderer: renderer.New(rendererType, palette),
	}
}

// AskMazeDimensions cпрашивает ширину и высоту.
func (c *console) AskMazeDimensions() (height, width int, err error) {
	areValids := func(data ...any) bool {
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

	c.askCorrectData(
		"%s\n",
		DIMENSIONS_INPUT_MESSAGE,
		ERROR_DIMENSIONS_INPUT_MESSAGE,
		areValids,
		&width, &height,
	)

	return height, width, nil
}

// AskCoordinates cпрашивает координаты start и end.
func (c *console) AskCoordinates(height, width int) (start, end cells.Coordinates, err error) {
	var x, y int

	areValids := func(data ...any) bool {
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

	c.printf("\n%s\n", NOTE_MESSAGE)

	c.askCorrectData(
		"%s\n",
		START_INPUT_MESSAGE,
		ERROR_COORDINATES_INPUT_MESSAGE,
		areValids,
		&x, &y,
	)

	start = cells.Coordinates{
		X: x,
		Y: y,
	}

	c.askCorrectData(
		"\n%s\n",
		END_INPUT_MESSAGE,
		ERROR_COORDINATES_INPUT_MESSAGE,
		areValids,
		&x, &y,
	)

	end = cells.Coordinates{
		X: x,
		Y: y,
	}

	return start, end, nil
}

// DisplayMaze отображает лабиринт.
func (c *console) DisplayMaze(maze maze.Maze) {
	c.printf("\n%s\n", c.renderer.Render(maze))
}

// DisplayMazeWithPath отображает лабиринт и путь на нём.
func (c *console) DisplayMazeWithPath(maze maze.Maze, path []cells.Coordinates) {
	c.printf("\n%s\n", c.renderer.RenderPath(maze, path))
}

// askCorrectData спрашивает данные до тех пор, пока они не будут корректными, читая их в data...;
// данные, которые нужно спросить, должны передаваться по указателю.
func (c *console) askCorrectData(
	messageFormat, message, errorMessage string,
	areValid func(data ...any) bool,
	data ...any,
) {
	c.printf(messageFormat, message)

	for {
		err := c.read(data...)
		if err != nil || !areValid(data...) {
			c.printf("%s", ERROR_COORDINATES_INPUT_MESSAGE)
		} else {
			break
		}
	}
}

// read читает данные в data.
func (c *console) read(data ...any) error {
	_, err := fmt.Fscan(&c.reader, data...)
	if err != nil {
		fmt.Fscanln(&c.reader)
		return fmt.Errorf("can`t scan data: %w", err)
	}

	return nil
}

// printf выводит данные согласно формату.
func (c *console) printf(format string, a ...any) error {
	_, err := fmt.Fprintf(&c.writer, format, a...)
	if err != nil {
		return fmt.Errorf("can`t write formatted data: %w", err)
	}

	err = c.writer.Flush()
	if err != nil {
		return fmt.Errorf("can`t flush data: %w", err)
	}

	return nil
}
