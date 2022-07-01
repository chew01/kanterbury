package proxy

type GameState struct {
	Player    *PlayerData
	Character *CharacterData
	Activity  *ActivityData
	PingFn    func(state *GameState)
}

func (gs *GameState) SetFn(fn func(state *GameState)) {
	gs.PingFn = fn
}

func (gs *GameState) Ping() {
	gs.PingFn(gs)
}
