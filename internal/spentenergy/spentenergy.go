package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive")
	}
	if duration <= 0 || weight <= 0 || height <= 0 || steps <= 0 {
		return 0, fmt.Errorf("invalid parameters")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive")
	}
	if duration <= 0 || weight <= 0 || height <= 0 || steps <= 0 {
		return 0, fmt.Errorf("invalid parameters")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	userDistance := Distance(steps, height)
	timeInHours := duration.Hours()
	return userDistance / timeInHours
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	userDistance := (height * stepLengthCoefficient * float64(steps)) / mInKm
	return userDistance
}
