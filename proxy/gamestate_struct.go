package proxy

type PlayerData struct {
	PlayerID  string `json:"playerId"`
	AppID     string `json:"appId"`
	GameWorld string `json:"gameWorld"`
	StartTime int64  `json:"modTime"`
}

type CharacterData struct {
	ID    int64  `json:"character1Id"`
	Level int64  `json:"character1Lv"` // Character lv -1
	Name  string `json:"character1Nm"`
}

type ActivityData struct {
	Name    string `json:"gameMode"` // For pvp
	EndTime int64  `json:"endTime"`
}
