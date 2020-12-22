package cgol

//reps the world, the game etc....
type Game struct {
	world *world
}

func NewGame(width, height int) *Game {
	game := &Game{
		world: newWorld(width, height),
	}
	return game
}
func (g *Game) Width() int {
	return g.world.width
}
func (g *Game) Height() int {
	return g.world.height
}

func (g *Game) Get(x, y int) bool {
	return g.world.Get(x, y)
}
func (g *Game) Set(x, y int, v bool) {
	g.world.Set(x, y, v)
}
func (g *Game) Step() {
	nextWorld := newWorld(g.world.width, g.world.height)
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			nextCell := nextCellState(g.world.Get(x, y), g.world.countLiveNeighbors(x, y))
			nextWorld.Set(x, y, nextCell)
		}
	}
	g.world = nextWorld
}

type world struct {
	cells         []bool
	width, height int
}

func newWorld(width, height int) *world {
	w := &world{
		cells:  make([]bool, width*height),
		width:  width,
		height: height,
	}
	return w
}

func (g *world) Get(x, y int) bool {
	idx := y*g.width + x
	return g.cells[idx]
}
func (g *world) safeGet(x, y int) bool {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return false
	}
	return g.Get(x, y)
}

func (g *world) countLiveNeighbors(x, y int) int {
	count := 0
	if g.safeGet(x-1, y-1) {
		count++
	}
	if g.safeGet(x, y-1) {
		count++
	}
	if g.safeGet(x+1, y-1) {
		count++
	}
	if g.safeGet(x-1, y) {
		count++
	}
	if g.safeGet(x+1, y) {
		count++
	}
	if g.safeGet(x-1, y+1) {
		count++
	}
	if g.safeGet(x, y+1) {
		count++
	}
	if g.safeGet(x+1, y+1) {
		count++
	}
	return count
}

func (g *world) Set(x, y int, v bool) {
	idx := y*g.width + x
	g.cells[idx] = v
}

func nextCellState(myLife bool, neighbors int) bool {
	return (myLife && neighbors == 2) || neighbors == 3
}
