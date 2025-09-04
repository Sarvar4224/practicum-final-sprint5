package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	newSlice := strings.Split(datastring, ",")
	if len(newSlice) != 2 {
		return fmt.Errorf("arguments must equal 2")
	}
	ds.Steps, err = strconv.Atoi(newSlice[0])
	if err != nil {
		return fmt.Errorf("invalid steps format %w", err)
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("steps must be positive")
	}
	ds.Duration, err = time.ParseDuration(newSlice[1])
	if err != nil {
		return fmt.Errorf("invalid timeDuration format: %w", err)
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("invalid duration")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("wrong format %w", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCalories), nil
}
