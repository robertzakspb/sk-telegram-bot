package main

import (
	"testing"
	"time"
)

func TestBreadPollOptionTitleGeneration(t *testing.T) {
	//Arrange
	randomSunday := time.Date(2025, 1, 5, 0,0,0,0, time.UTC)
	expectedSundayTitle := "°ХЛЕБ (1-2)°: из ДЦ \"Кров\""

	randomWednesday := time.Date(2025, 1, 1, 0,0,0,0, time.UTC)
	randomSaturday := time.Date(2025, 1, 4, 0,0,0,0, time.UTC)
	expectedSaturdayOrWednesdayTitle := "°ХЛЕБ (1-2)°: из продавнице \"Расина\""
	
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

	if wednesdayTitle != expectedSaturdayOrWednesdayTitle {
		t.Fatalf("Expected Wednesday title: %v; actual: %v",expectedSaturdayOrWednesdayTitle, wednesdayTitle )
	}

	if saturdayTitle != expectedSaturdayOrWednesdayTitle {
		t.Fatalf("Expected Saturday title: %v; actual: %v",expectedSaturdayOrWednesdayTitle, saturdayTitle )
	}

	if sundayTitle != expectedSundayTitle {
		t.Fatalf("Expected Sunday title: %v; actual: %v",expectedSundayTitle, sundayTitle )
	}
}