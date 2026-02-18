package spentenergy

import (
	"errors"
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
		err := errors.New("Количество шагов должно быть больше 0")
		return 0, err

	}
	if weight <= 0 {
		err := errors.New("Вес должен быть больше 0")
		return 0, err

	}

	if height <= 0 {
		err := errors.New("Рост должен быть больше 0")
		return 0, err

	}

	if duration <= 0 {
		err := errors.New("Длительность должна быть больше 0")
		return 0, err

	}
	averspeed := MeanSpeed(steps, height, duration)
	return (walkingCaloriesCoefficient * weight * averspeed * duration.Minutes()) / minInH, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		err := errors.New("Количество шагов должно быть больше 0")
		return 0, err

	}
	if weight <= 0 {
		err := errors.New("Вес должен быть больше 0")
		return 0, err

	}

	if height <= 0 {
		err := errors.New("Рост должен быть больше 0")
		return 0, err

	}

	if duration <= 0 {
		err := errors.New("Длительность должна быть больше 0")
		return 0, err

	}

	averspeed := MeanSpeed(steps, height, duration)
	return (weight * averspeed * duration.Minutes()) / minInH, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dis := Distance(steps, height)
	avespeed := dis / duration.Hours()
	return avespeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	steplen := height * stepLengthCoefficient
	return (float64(steps) * steplen) / mInKm
}
