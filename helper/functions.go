package helper

import (
	"fmt"
	"groupie-tracker/models"
	"os"
	"strings"
	"time"
)

func FormatDate(dateStr string) string {
	parsedDate, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		os.Exit(1)
	}

	formatted := parsedDate.Format("02 Jan 2006")
	return formatted
}

func FormatConcertDates(relations *models.Relations) {
	for key, dates := range relations.DatesLocations {
		for y, date := range dates {
			relations.DatesLocations[key][y] = FormatDate(date)
		}
	}
}

func FormatLocations(locations *models.Location) {
	for i, location := range locations.Locations {
		locations.Locations[i] = strings.ReplaceAll(location, "_", " ")
		locations.Locations[i] = strings.ReplaceAll(locations.Locations[i], "-", ",")
		locations.Locations[i] = Capitalize(locations.Locations[i])

	}

}


func AppendIfNotExist(slice []string, value string) []string {
	for _, v := range slice {
		if v == value {
			return slice
		}
	}
	return append(slice, value)
}

func Capitalize(s string) string {
	runes := []rune(s)
	var check = func(a rune) bool {
		if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
			return true
		}
		return false
	}
	first := true
	for i := range runes {
		if check(runes[i]) && first {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] -= 32
			}
			first = false
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		} else if !check(runes[i]) {
			first = true
		}
	}
	return string(runes)
}

func HasKeyword(keyword string, s string) bool{
	words := strings.Split(s,",")

	if strings.Contains(keyword,",")  {
		return strings.HasPrefix(s,keyword)
	}

	for _, v := range words {
			if strings.HasPrefix(v,keyword){
				return true
			}
	}
	return false
}

