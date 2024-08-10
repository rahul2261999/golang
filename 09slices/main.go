package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the slices")

	var friuitList = []string{"apple", "cherry", "date"}

	fmt.Printf("Type of friuitList: %T\n", friuitList)

	friuitList = append(friuitList, "Mango")

	fmt.Println("Updated list of fruits:", friuitList)

	friuitList = append(friuitList[1:], "banana")

	fmt.Println("Updated list of fruits:", friuitList)

	highScores := make([]int, 5)

	highScores[0] = 90
	highScores[1] = 85
	highScores[2] = 80
	highScores[3] = 75
	highScores[4] = 70

	highScores = append(highScores, 80, 95, 100)

	fmt.Println("High scores:", highScores)

	sort.Ints(highScores)

	fmt.Println("Sorted High scores:", highScores)

	var courses = []string{"Go", "Python", "JavaScript", "C++", "Rust"}
	fmt.Println("Courses:", courses)

	const removeIndex = 2
	courses = append(courses[:removeIndex], courses[removeIndex+1:]...)
	fmt.Println("Courses:", courses)
	fmt.Println("Courses len:", len(courses))

}
