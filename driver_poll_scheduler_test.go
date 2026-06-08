package main

import (
	"testing"
	"time"
)

func Test_generateDriverPollTitle_Monday(t *testing.T) {
	monday := time.Date(2026, 6, 8, 0, 0, 0, 0, time.UTC)
	expectedTitle := "8.6 – 14.6"

	title := generateDriverPollTitle(monday)

	if title != expectedTitle {
		t.Fatalf("Expected: %v, got: %v", expectedTitle, title)
	}
}

func Test_generateDriverPollTitle_Sunday(t *testing.T) {
	sunday := time.Date(2026, 6, 7, 0, 0, 0, 0, time.UTC)
	expectedTitle := "8.6 – 14.6"

	title := generateDriverPollTitle(sunday)

	if title != expectedTitle {
		t.Fatalf("Expected: %v, got: %v", expectedTitle, title)
	}
}

func Test_getThisOrNextMonday_OnSunday(t *testing.T) {
	sunday := time.Date(2026, 6, 7, 0, 0, 0, 0, time.UTC)

	year, month, date := getThisOrNextMonday(sunday).Date()

	if !(date == 8 && month == 6 && year == 2026) {
		t.Fatalf("Expected date: 8.6.2026; actual date: %v", getThisOrNextMonday(sunday))
	}
}

func Test_getThisOrNextMonday_OnSaturday(t *testing.T) {
	saturday := time.Date(2026, 6, 6, 0, 0, 0, 0, time.UTC)

	year, month, date := getThisOrNextMonday(saturday).Date()

	if !(date == 1 && month == 6 && year == 2026) {
		t.Fatalf("Expected date: 8.6.2026; actual date: %v", getThisOrNextMonday(saturday))
	}
}

func Test_getThisOrNextMonday_OnMonday(t *testing.T) {
	monday := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)

	year, month, date := getThisOrNextMonday(monday).Date()

	if !(date == 1 && month == 6 && year == 2026) {
		t.Fatalf("Expected date: 8.6.2026; actual date: %v", getThisOrNextMonday(monday))
	}
}

func Test_getThisOrNextMonday_OnWednesday(t *testing.T) {
	wednesday := time.Date(2026, 6, 3, 0, 0, 0, 0, time.UTC)

	year, month, date := getThisOrNextMonday(wednesday).Date()

	if !(date == 1 && month == 6 && year == 2026) {
		t.Fatalf("Expected date: 8.6.2026; actual date: %v", getThisOrNextMonday(wednesday))
	}
}
