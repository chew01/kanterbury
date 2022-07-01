package utils

import "strings"

var worldMap = map[string]string{
	"app_id_146318_world_1": "Asia 1",
	"app_id_146318_world_2": "Asia 2",
}

func FormatName(raw string) string {
	rawArr := strings.Split(raw, "_")
	formattedArr := make([]string, len(rawArr))

	for _, word := range rawArr {
		formattedWord := strings.ToUpper(word[0:1]) + word[1:]
		formattedArr = append(formattedArr, formattedWord)
	}
	return strings.Join(formattedArr, " ")
}

func FormatLevel(raw int64) int64 {
	return raw + 1
}

func FormatWorld(raw string) string {
	return worldMap[raw]
}
