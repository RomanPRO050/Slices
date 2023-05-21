package main

import "fmt"

var workDays = []string{"пон", "вт", "ср", "чт", "пт"}
var weekendDays = []string{"сб", "вс"}
var weekDays = make([]string, 0, 7)

func main() {
	weekDays = append(workDays, weekendDays[0], weekendDays[1])
	fmt.Println(weekDays)
}
