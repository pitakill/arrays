package main

import (
	"fmt"
	"strconv"
	"strings"
)

func transformToString(input []int) string {
	output := make([]string, 0)

	for _, e := range input {
		output = append(output, strconv.Itoa(e))
	}

	return fmt.Sprintf("[%s]", strings.Join(output, " "))
}

func printWithColors(data interface{}, color int) string {
	var dataCasted string

	switch casted := data.(type) {
	case []int:
		dataCasted = transformToString(casted)
	case int:
		dataCasted = strconv.Itoa(casted)
	case string:
		dataCasted = casted
	}

	return fmt.Sprintf("\033[38;5;%dm%s\033[39;49m", color, dataCasted)
}
