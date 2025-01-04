package main

import (
	"testing"
	"time"
)

func Test_BreadPollOptionTitleGeneration(t *testing.T) {
	//Arrange
	randomSunday := time.Date(2025, 1, 5, 0, 0, 0, 0, time.UTC)
	expectedSundayTitle := "°ХЛЕБ (1-2)°: из ДЦ \"Кров\""

	randomWednesday := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	randomSaturday := time.Date(2025, 1, 4, 0, 0, 0, 0, time.UTC)
	expectedSaturdayOrWednesdayTitle := "°ХЛЕБ (1-2)°: из продавнице \"Расина\""

	randomMonday := time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC)
	expectedNonDistributionDateTitle := "°ХЛЕБ (1-2)°"

	//Act
	wednesdayTitle := breadPollOptionTitle(randomWednesday)
	saturdayTitle := breadPollOptionTitle(randomSaturday)
	sundayTitle := breadPollOptionTitle(randomSunday)
	nonDistributionDateTitle := breadPollOptionTitle(randomMonday)

	//Assert
	if nonDistributionDateTitle != expectedNonDistributionDateTitle {
		t.Fatalf("Expected title for a non-distribution date: %v; actual: %v", expectedNonDistributionDateTitle, nonDistributionDateTitle)
	}

	if wednesdayTitle != expectedSaturdayOrWednesdayTitle {
		t.Fatalf("Expected Wednesday title: %v; actual: %v", expectedSaturdayOrWednesdayTitle, wednesdayTitle)
	}

	if saturdayTitle != expectedSaturdayOrWednesdayTitle {
		t.Fatalf("Expected Saturday title: %v; actual: %v", expectedSaturdayOrWednesdayTitle, saturdayTitle)
	}

	if sundayTitle != expectedSundayTitle {
		t.Fatalf("Expected Sunday title: %v; actual: %v", expectedSundayTitle, sundayTitle)
	}
}

//Tests if the number of required people on different dates is properly calculated
func Test_RequiredNumberOfPeople(t *testing.T) {
	//Arrange
	randomSunday := time.Date(2025, 1, 5, 0, 0, 0, 0, time.UTC)
	expectedNumberOfPeopleOnSunday := "(6-7)"
	actualNumberOfPeopleOnSunday := requiredNumberOfPeople(randomSunday)

	randomWednesday := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedNumberOfPeopleOnWednesday := "(5)"
	actualNumberOfPeopleOnWednesday := requiredNumberOfPeople(randomWednesday)

	//Assert
	if actualNumberOfPeopleOnSunday != expectedNumberOfPeopleOnSunday {
		t.Fatalf("Expected number of people on Sunday: %v; actual: %v", expectedNumberOfPeopleOnSunday, actualNumberOfPeopleOnSunday)
	}

	if actualNumberOfPeopleOnWednesday != expectedNumberOfPeopleOnWednesday {
		t.Fatalf("Expected number of people on Wednesday: %v; actual: %v", expectedNumberOfPeopleOnWednesday, actualNumberOfPeopleOnWednesday)
	}
}

//Tests if the drink distribution day (Sunday) is properly determined
func Test_ShouldDistributeDrinks(t * testing.T) {
	//Arrange
	randomSunday := time.Date(2025, 1, 5, 0, 0, 0, 0, time.UTC)
	randomWednesday := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	//Assert
	if shouldDistributeDrink(randomSunday) != true {
		t.Fatalf("Expected: drink should be distributed on Sunday (%v)", randomSunday)
	}

	if shouldDistributeDrink(randomWednesday) != false {
		t.Fatalf("Expected: drink should not be distributed on Wednesday (%v)", randomWednesday)
	}
}

func Test_MealDistributionPollTitle(t * testing.T) {
	//Arrange
	randomWednesday := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedWednesdayPollTitle := "СРЕДА 1.1.2025"

	randomSaturday := time.Date(2025, 1, 4, 0, 0, 0, 0, time.UTC)
	expectedSaturdayPollTitle := "СУБОТА 4.1.2025"

	randomSunday := time.Date(2025, 1, 5, 0, 0, 0, 0, time.UTC)
	expectedSundayPollTitle := "НЕДЕЉА 5.1.2025"

	//Assert
	if generateMealDistributionPollTitle(randomWednesday) != expectedWednesdayPollTitle {
		t.Fatalf("Incorrect poll title. Expected: %v; actual: %v", expectedWednesdayPollTitle, generateMealDistributionPollTitle(randomWednesday))
	}

	//Assert
	if generateMealDistributionPollTitle(randomSaturday) != expectedSaturdayPollTitle {
		t.Fatalf("Incorrect poll title. Expected: %v; actual: %v", expectedSaturdayPollTitle, generateMealDistributionPollTitle(randomSaturday))
	}

	if generateMealDistributionPollTitle(randomSunday) != expectedSundayPollTitle {
		t.Fatalf("Incorrect poll title. Expected: %v; actual: %v", expectedSundayPollTitle, generateMealDistributionPollTitle(randomSunday))
	}
}