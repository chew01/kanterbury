package proxy

type GameState struct {
	Player    *PlayerData
	Character *CharacterData
	Activity  *ActivityData
	Hook      chan *GameState
}

func (gs *GameState) SetHook(hook chan *GameState) {
	gs.Hook = hook
}

func (gs *GameState) Ping() {
	if gs.Hook != nil {
		gs.Hook <- gs
	}
}
