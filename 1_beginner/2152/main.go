package main

import (
	"fmt"
	"strconv"
)

func convert(value int) string {
	if value < 10 {
		return "0" + strconv.Itoa(value)
	}
	return strconv.Itoa(value)
}

func doorText(value int) string {
	if value == 1 {
		return "A porta abriu!"
	}
	return "A porta fechou!"
}

func main() {
	var total, hours, minutes, door int
	fmt.Scanf("%d", &total)
	for total > 0 {
		fmt.Scanf("%d %d %d", &hours, &minutes, &door)
		fmt.Println(convert(hours) + ":" + convert(minutes) + " - " + doorText(door))
		total--
	}
}
