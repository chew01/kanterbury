package proxy

import (
	"errors"
	"github.com/chew01/kanterbury/utils"
	"time"
)

type GameState struct {
	Player    *PlayerData
	Startup   *StartupData
	Character *CharacterData
	Activity  *ActivityData
	Hooks     map[string]func(state *GameState)
}

func newGameState() *GameState {
	return &GameState{
		Player:    &PlayerData{},
		Startup:   &StartupData{StartTime: time.Now().Unix()},
		Character: &CharacterData{},
		Activity:  &ActivityData{},
		Hooks:     make(map[string]func(state *GameState)),
	}
}

// AddHook adds a hook function that will be executed with the game state pointer when the state is updated
func (gs *GameState) AddHook(key string, fn func(state *GameState)) error {
	if gs.Hooks[key] != nil {
		return errors.New("failed to add hook to game state: already contains key " + key)
	}
	gs.Hooks[key] = fn
	return nil
}

// RemoveHook removes an existing hook function with the given key
func (gs *GameState) RemoveHook(key string) error {
	if gs.Hooks[key] == nil {
		return errors.New("failed to remove hook to game state: does not contain key " + key)
	}
	gs.Hooks[key] = nil
	return nil
}

func (gs *GameState) updatePlayer(newData *PlayerData) {
	newData.GameWorld = utils.FormatWorld(newData.GameWorld)
	gs.Player = newData
	gs.pingHooks()
}

func (gs *GameState) updateStartup(newData *StartupData) {
	gs.Startup = newData
	gs.pingHooks()
}

func (gs *GameState) updateCharacter(newData *CharacterData) {
	newData.Name = utils.FormatName(newData.Name)
	newData.Level = utils.FormatLevel(newData.Level)
	gs.Character = newData
	gs.pingHooks()
}

func (gs *GameState) updateActivity(newData *ActivityData) {
	gs.Activity = newData
	gs.pingHooks()
}

func (gs *GameState) pingHooks() {
	for _, fn := range gs.Hooks {
		fn(gs)
	}
}
