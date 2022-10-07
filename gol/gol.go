package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	for i := 0; i < p.imageWidth; i++ {
		for j := 0; i < p.imageHeight; j++ {
			var aliveNeighbours = getLiveNeighbours(p, world, i, j)
			if isAlive(world[i][j]) {
				if aliveNeighbours < 4 {
					if aliveNeighbours > 1 {
						// pass really doesn't exist in golang????????
					} else {
						world[i][j] = 0
					}
				} else {
					world[i][j] = 0
				}
			} else {
				if aliveNeighbours == 3 {
					world[i][j] = 255
				}
			}
		}
	}
	return world
}

func getLiveNeighbours(p golParams, world [][]byte, a, b int) int {
	var alive = 0
	var widthLeft int
	var widthRight int
	var heightUp int
	var heightDown int

	if a == 0 {
		widthLeft = p.imageWidth
	} else {
		widthLeft = a - 1
	}
	if a == p.imageWidth {
		widthRight = 0
	} else {
		widthRight = a + 1
	}

	if b == 0 {
		heightUp = p.imageHeight
	} else {
		heightUp = b + 1
	}

	if b == p.imageHeight {
		heightDown = 0
	} else {
		heightUp = b - 1
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
		for j := 0; i < p.imageHeight; j++ {
			if isAlive(world[i][j]) {
				cells = append(cells, cell{i, j})
			}
		}
	}

	return []cell{}
}
