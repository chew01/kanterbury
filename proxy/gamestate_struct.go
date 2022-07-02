package proxy

// PlayerData for data related to the player. Updated on every login/resource check
type PlayerData struct {
	PlayerID  string `json:"playerId"`
	AppID     string `json:"appId"`
	GameWorld string `json:"gameWorld"`
}

// StartupData for setting the initial login time. Updated on every login
type StartupData struct {
	StartTime int64 `json:"modTime"`
}

// CharacterData for data related to the first character of the first team. Updated on every resource check
type CharacterData struct {
	ID    int64  `json:"character1Id"`
	Level int64  `json:"character1Lv"` // Character lv -1
	Name  string `json:"character1Nm"`
}

// ActivityData for multiplayer data (specifically PvP). Updated on every multiplayer gamemode round
type ActivityData struct {
	Name    string `json:"gameMode"` // For pvp
	EndTime int64  `json:"endTime"`
}
