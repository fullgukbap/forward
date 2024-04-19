package netcode

// Player represents a player in theEngine
type Player struct {
	Name string
	X    int
	Y    int
}

// Engine represents the current state of theEngine
type Engine struct {
	Players []*Player
}

func NewEngine() *Engine {
	return &Engine{Players: make([]*Player)}
}

func (e *Engine) AddPlayer(player *Player) {
	e.Players = append(e.Players, player)
}

func (e *Engine) RemovePlayer(n string) {
	findIdx := -1
	for i, p := range e.Players {
		if p.Name == n {
			findIdx = i
			break	
		}
	}
	if findIdx 
}

func (g *Engine) MovePlayer(playerName string, deltaX, deltaY int) {
	player := g.Players[playerName]
	player.X += deltaX
	player.Y += deltaY
}
// 비번 너무 쉽네요