package main

import "fmt"

var weekDays = []string{"пон", "вт", "ср", "чт", "пт", "сб", "вс"}
var workDays = make([]string, 5, 7)
var weekendDays = make([]string, 2, 7)

func main() {
	//workDays = weekDays[0:5]
	copy(workDays, weekDays[0:5])
	copy(weekendDays, weekDays[5:7])
	fmt.Println(workDays, weekendDays)
}
