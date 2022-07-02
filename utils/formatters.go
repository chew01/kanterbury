package utils

import "strings"

// Mapping of each game world string to their formatted version
var worldMap = map[string]string{
	"app_id_146318_world_1": "Asia 1",
	"app_id_146318_world_2": "Asia 2",
}

// FormatName formats the raw name of a character in the form "example_name" to "Example Name"
func FormatName(raw string) string {
	rawArr := strings.Split(raw, "_")
	formattedArr := make([]string, len(rawArr))

	for _, word := range rawArr {
		formattedWord := strings.ToUpper(word[0:1]) + word[1:]
		formattedArr = append(formattedArr, formattedWord)
	}
	return strings.Join(formattedArr, " ")
}

// FormatLevel formats the raw level of a character to their actual level
func FormatLevel(raw int64) int64 {
	return raw + 1
}

// FormatWorld formats the raw world string of a player to their actual world name
func FormatWorld(raw string) string {
	return worldMap[raw]
}
