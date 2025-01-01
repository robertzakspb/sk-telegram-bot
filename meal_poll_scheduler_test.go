package main

import (
	"testing"
	"time"
)

func TestBreadPollOptionTitleGeneration(t *testing.T) {
	//Arrange
	randomWednesday := time.Date(2025, 1, 1, 0,0,0,0, time.UTC)
	randomSunday := time.Date(2025, 1, 5, 0,0,0,0, time.UTC)
	expectedWednesdayOrSundayTitle := "°ХЛЕБ (1-2)°: из ДЦ \"Кров\""

	randomSaturday := time.Date(2025, 1, 4, 0,0,0,0, time.UTC)
	expectedSaturdayTitle := "°ХЛЕБ (1-2)°: из продавнице \"Расина\""
	
	//Test for a non-distribution date
	randomMonday := time.Date(2024, 12, 30, 0,0,0,0, time.UTC)
	expectedNonDistributionDateTitle := "°ХЛЕБ (1-2)°"


	//Act
	wednesdayTitle := breadPollOptionTitle(randomWednesday)
	saturdayTitle := breadPollOptionTitle(randomSaturday)
	sundayTitle := breadPollOptionTitle(randomSunday)
	nonDistributionDateTitle := breadPollOptionTitle(randomMonday)

	//Assert 
	if nonDistributionDateTitle !=  expectedNonDistributionDateTitle {
		t.Fatalf("Expected title for a non-distribution date: %v; actual: %v",expectedNonDistributionDateTitle, nonDistributionDateTitle )
	}

	if wednesdayTitle != expectedWednesdayOrSundayTitle {
		t.Fatalf("Expected Wednesday title: %v; actual: %v",expectedWednesdayOrSundayTitle, wednesdayTitle )
	}

	if saturdayTitle != expectedSaturdayTitle {
		t.Fatalf("Expected Saturday title: %v; actual: %v",expectedWednesdayOrSundayTitle, saturdayTitle )
	}

	if sundayTitle != expectedWednesdayOrSundayTitle {
		t.Fatalf("Expected Sunday title: %v; actual: %v",expectedWednesdayOrSundayTitle, sundayTitle )
	}
}