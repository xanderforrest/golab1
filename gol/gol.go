package main

func calculateNextState(p golParams, world [][]byte) [][]byte {

	newWorld := make([][]uint8, p.imageWidth)
	for i := range newWorld {
		newWorld[i] = make([]uint8, p.imageHeight)
	}

	for i := 0; i < p.imageWidth; i++ {
		for j := 0; j < p.imageHeight; j++ {
			var aliveNeighbours = getLiveNeighbours(p, world, i, j)
			if isAlive(world[i][j]) {
				if aliveNeighbours < 4 {
					if aliveNeighbours > 1 {
						newWorld[i][j] = world[i][j]
						// pass really doesn't exist in golang????????
					} else {
						newWorld[i][j] = 0
					}
				} else {
					newWorld[i][j] = 0
				}
			} else {
				if aliveNeighbours == 3 {
					newWorld[i][j] = 255
				} else {
					newWorld[i][j] = 0
				}
			}
		}
	}
	return newWorld
}

func getLiveNeighbours(p golParams, world [][]byte, a, b int) int {
	var alive = 0
	var widthLeft int
	var widthRight int
	var heightUp int
	var heightDown int

	if a == 0 {
		widthLeft = p.imageWidth - 1
	} else {
		widthLeft = a - 1
	}
	if a == p.imageWidth-1 {
		widthRight = 0
	} else {
		widthRight = a + 1
	}

	if b == 0 {
		heightUp = p.imageHeight - 1
	} else {
		heightUp = b - 1
	}

	if b == p.imageHeight-1 {
		heightDown = 0
	} else {
		heightDown = b + 1
	}

	if isAlive(world[widthLeft][b]) {
		alive = alive + 1
	}
	if isAlive(world[widthRight][b]) {
		alive = alive + 1
	}
	if isAlive(world[widthLeft][heightUp]) {
		alive = alive + 1
	}
	if isAlive(world[a][heightUp]) {
		alive = alive + 1
	}
	if isAlive(world[widthRight][heightUp]) {
		alive = alive + 1
	}
	if isAlive(world[widthLeft][heightDown]) {
		alive = alive + 1
	}
	if isAlive(world[a][heightDown]) {
		alive = alive + 1
	}
	if isAlive(world[widthRight][heightDown]) {
		alive = alive + 1
	}
	return alive
}

func isAlive(cell byte) bool {
	if cell == 255 {
		return true
	}
	return false
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var cells []cell

	for i := 0; i < p.imageWidth; i++ {
		for j := 0; j < p.imageHeight; j++ {
			if isAlive(world[i][j]) {
				cells = append(cells, cell{j, i})
			}
		}
	}

	return cells
}
