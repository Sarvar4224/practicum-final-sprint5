package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	newSlice := strings.Split(datastring, ",")
	if len(newSlice) != 3 {
		return fmt.Errorf("arguments must equal 3")
	}
	t.Steps, err = strconv.Atoi(newSlice[0])
	if err != nil {
		return fmt.Errorf("invalid steps format %w", err)
	}
	if t.Steps <= 0 {
		return fmt.Errorf("steps must be positive")
	}

	t.TrainingType = newSlice[1]

	t.Duration, err = time.ParseDuration(newSlice[2])
	if err != nil {
		return fmt.Errorf("invalid timeDuration format: %w", err)
	}
	if t.Duration <= 0 {
		return fmt.Errorf("invalid duration")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var spentCalories float64

	switch t.TrainingType {
	case "Бег":
		spentCalories, _ = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	case "Ходьба":
		spentCalories, _ = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, spentCalories)

	return result, nil
}
