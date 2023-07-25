package main

import (
	"strings"
	"unicode"

	"github.com/go-rod/rod"
	"mvdan.cc/xurls/v2"
)

// ClassInformation holds the class details scraped from the webpage
type ClassInformation struct {
	Course     string `json:"course"`
	Title      string `json:"title"`
	Name       string `json:"name"`
	Class      string `json:"class"`
	Instructor string `json:"instructor"`
	Days       string `json:"days"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Location   string `json:"location"`
	Dates      string `json:"dates"`
	Units      string `json:"units"`
	Seats      string `json:"seats"`
	Syllabus   string `json:"syllabus"`
	GS         string `json:"gs"`
	IsOpen     bool   `json:"is_open"`
}

// Returns information about a specific class by ID
func scrapeClass(classId string) []ClassInformation {
	// If class ID is invalid, return an empty list
	if len(classId) != 5 {
		return []ClassInformation{}
	}
	return scrape("https://catalog.apps.asu.edu/catalog/classes/classlist?campusOrOnlineSelection=C&honors=F&keywords=" + classId + "&promod=F&searchType=all&term=2237")
}

// Returns information about all classes
func scrapeClasses() []ClassInformation {
	return scrape("https://catalog.apps.asu.edu/catalog/classes/classlist?campusOrOnlineSelection=C&honors=F&promod=F&searchType=all&term=2237")
}

// Returns information about all classes for a specific course
func scrapeClassesByCourse(subject string, number string) []ClassInformation {
	// If subject and number are invalid, return an empty list
	if len(subject) != 3 || len(number) != 3 || !IsLetter(subject) || !IsLetter(number) {
		return []ClassInformation{}
	}

	return scrape("https://catalog.apps.asu.edu/catalog/classes/classlist?campusOrOnlineSelection=C&catalogNbr=" + number + "&honors=F&promod=F&searchType=all&subject=" + subject + "&term=2237")
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Helper function to scrape classes from the given URL
func scrape(url string) []ClassInformation {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(url).MustWaitStable()

	rows := page.MustElement(".class-results-rows").MustElements(".class-accordion")

	courses := make([]ClassInformation, 0, len(rows))

	for _, element := range rows {
		classInfo := element.MustElements(".class-results-cell")

		// Filter out elements with 'Multiple dates and times' text
		filteredClassInfo := make([]*rod.Element, 0, len(classInfo))
		for _, elem := range classInfo {
			if elem.MustText() != "Multiple dates and times" {
				filteredClassInfo = append(filteredClassInfo, elem)
			}
		}

		// Ensure that we have the correct number of data points for a class
		if len(filteredClassInfo) != 15 {
			continue
		}

		// Determine if there are open seats in the class
		openSeat := strings.TrimSpace(filteredClassInfo[11].MustText())[0] != '0'

		// Extract link to the class syllabus
		Link := ""
		xurlsStrict := xurls.Strict()
		if matches := xurlsStrict.FindAllString(strings.TrimSpace(filteredClassInfo[12].MustHTML()), -1); len(matches) > 0 {
			Link = matches[0]
		}

		// Create and append new ClassInformation
		course := ClassInformation{
			Course:     strings.TrimSpace(filteredClassInfo[0].MustText()),
			Title:      strings.TrimSpace(filteredClassInfo[1].MustText()),
			Name:       strings.TrimSpace(filteredClassInfo[2].MustText()),
			Class:      strings.TrimSpace(filteredClassInfo[3].MustText()),
			Instructor: strings.TrimSpace(filteredClassInfo[4].MustText()),
			Days:       strings.TrimSpace(filteredClassInfo[5].MustText()),
			StartTime:  strings.TrimSpace(filteredClassInfo[6].MustText()),
			EndTime:    strings.TrimSpace(filteredClassInfo[7].MustText()),
			Location:   strings.ReplaceAll(strings.TrimSpace(filteredClassInfo[8].MustText()), "\n", " "),
			Dates:      strings.ReplaceAll(strings.TrimSpace(filteredClassInfo[9].MustText()), "\n", " "),
			Units:      strings.TrimSpace(filteredClassInfo[10].MustText()),
			Seats:      strings.TrimSpace(filteredClassInfo[11].MustText()),
			Syllabus:   Link,
			GS:         strings.TrimSpace(filteredClassInfo[13].MustText()),
			IsOpen:     openSeat,
		}

		courses = append(courses, course)
	}

	return courses
}
