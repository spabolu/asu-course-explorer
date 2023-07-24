package main

import (
	"errors"
	"fmt"
	"github.com/go-rod/rod"
	"strconv"
	"strings"
	"unicode"
)

type ClassInformation struct {
	Course     string `json:"course"`
	Title      string `json:"title"`
	Class      string `json:"class"`
	Instructor string `json:"instructor"`
	Days       string `json:"days"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Location   string `json:"location"`
	Dates      string `json:"dates"`
	Units      string `json:"units"`
	Seats      string `json:"seats"`
	IsOpen     bool   `json:"is_open"`
}

func scrapeClass(classNumber string) (ClassInformation, error) {
	if len(classNumber) != 5 {
		return ClassInformation{}, errors.New("class number must be 5 digits")
	}

	url := fmt.Sprintf("https://catalog.apps.asu.edu/catalog/classes/classlist?campusOrOnlineSelection=A&honors=F&keywords=%s&promod=F&searchType=all&term=2237", classNumber)

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(url).MustWaitStable()

	list := page.MustWaitStable().MustElements(".class-results-cell")

	var htmlDOM []string
	builder := &strings.Builder{}

	for _, element := range list {
		builder.Reset()
		text := strings.TrimRightFunc(element.MustText(), unicode.IsSpace)
		text = strings.ReplaceAll(text, "\n", " ")

		if text != "" {
			htmlDOM = append(htmlDOM, text)
		}
	}

	// TODO: Fix issue with iCourse, and multiple instructors courses
	removeIndices := map[int]struct{}{2: {}, 5: {}, 13: {}, 14: {}, 15: {}}

	if !strings.Contains(htmlDOM[9], "Hybrid") && !strings.Contains(htmlDOM[5], "Multiple") {
		removeIndices = map[int]struct{}{2: {}, 12: {}, 13: {}, 14: {}}
	}

	var classData []string
	for i, val := range htmlDOM {
		if _, found := removeIndices[i]; !found {
			classData = append(classData, val)
		}
	}

	classDetails := ClassInformation{
		Course:     getOrDefault(classData, 0),
		Title:      getOrDefault(classData, 1),
		Class:      getOrDefault(classData, 2),
		Instructor: parseInstructor(getOrDefault(classData, 3)),
		Days:       getOrDefault(classData, 4),
		StartTime:  getOrDefault(classData, 5),
		EndTime:    getOrDefault(classData, 6),
		Location:   getOrDefault(classData, 7),
		Dates:      getOrDefault(classData, 8),
		Units:      getOrDefault(classData, 9),
		Seats:      getOrDefault(classData, 10),
		IsOpen:     hasOpenSeats(getOrDefault(classData, 10)),
	}

	return classDetails, nil
}

func getOrDefault(data []string, index int) string {
	if index < len(data) {
		return data[index]
	}
	return ""
}

// TODO: Simplify the logic here
func hasOpenSeats(input string) bool {
	parts := strings.Split(input, " of ")
	if len(parts) != 2 {
		fmt.Println("Invalid input format")
		return false
	}

	firstNumber, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		fmt.Println("Error converting first part to integer:", err)
		return false
	}

	secondNumber, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		fmt.Println("Error converting second part to integer:", err)
		return false
	}

	return firstNumber <= secondNumber
}

func parseInstructor(instructor string) string {
	if strings.Contains(instructor, "Multiple instructors") {
		return "Multiple Instructors"
	}
	return instructor
}
