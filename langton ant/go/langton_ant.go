package main

import "fmt"

type Ant struct {
	x         int
	y         int
	direction string
}

func initializeMatrix() [][]int {
	matrix := make([][]int, 10)

	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			matrix[row] = append(matrix[row], 0)
		}
	}

	return matrix

}

func displayAnt(matrix [][]int) [][]int {
	ant := &Ant{
		antPositionx: 2,
		antPositiony: 2,
		antDirection: 0,
	}

	return matrix

}

var directions = []string{
	"north",
	"east",
	"south",
	"west",
}

var transitions = map[string][]string{
	"up":    []string{"right", "left"},
	"right": []string{"down", "up"},
	"down":  []string{"right", "left"},
	"left":  []string{"up", "down"},
}

type transition struct {
	dx, dy int
}

type position struct {
}

func (p *position) plus(t transition) *position {

}

type matrix struct {
}

func (a *Ant) move(t transition) *Ant {
	return &Ant{
		x: a.x + t.dx,
	}
}

var movements = map[string][]int{
	"up":    []int{0, -1},
	"down":  []int{0, 1},
	"left":  []int{-1, 0},
	"right": []int{1, 0},
}

func turnAnt(matrix [][]int, ant *Ant) {
	newDirectionRight := directions[(arrays.indexOf("east")+1)%4]
	newDirectionLeft := directions[(arrays.indexOf("east")-1)%4]

	// 0 white 1 black
	boxColor := matrix[ant.y][ant.x]

	if boxColor == 1 {
		matrix[ant.y][ant.x] = 0
	} else {
		matrix[ant.y][ant.x] = 1
	}

	ant.direction = transitions[ant.direction][boxColor]

	m := movements[ant.direction]
	ant.x += m[0]
	ant.y += m[1]
}
func main() {
	fmt.Println(displayAnt(initializeMatrix()))
}
