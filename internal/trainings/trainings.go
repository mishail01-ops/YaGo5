package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	//Steps int — количество шагов, проделанных за тренировку.
	//TrainingType string — тип тренировки(бег или ходьба).
	//Duration time.Duration — длительность тренировки.
	//personaldata.Personal — встроенная структура Personal из пакета personaldata, у которой есть метод Print().

	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	spl := strings.Split(datastring, ",")
	if len(spl) != 3 {
		err := errors.New("Данные неправильного формата, должно 3456,Ходьба,3h00m")
		return err
	}
	step, err := strconv.Atoi(spl[0])

	if err != nil {
		return err
	}

	if step <= 0 {
		err := errors.New("Количество шагов должно быть больше 0")
		return err
	}

	t.Steps = step
	t.TrainingType = spl[1]

	dur, err := time.ParseDuration(spl[2])

	if err != nil {
		return err
	}

	if dur <= 0 {
		err := errors.New("Длительность должна быть больше 0")
		return err
	}

	t.Duration = dur
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	var cal float64
	var err error
	err = nil
	dist := spentenergy.Distance(t.Steps, t.Height)
	averspeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	tp := t.TrainingType

	switch tp {
	case "Бег":
		{
			cal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		}

	case "Ходьба":
		{
			cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		}
	default:
		{
			err = errors.New("неизвестный тип тренировки")
		}
	}

	if err != nil {
		return "", err
	}

	dur := fmt.Sprintf("%d.%02d", int(t.Duration.Hours()), int(t.Duration.Minutes())%60)

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %s ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", tp, dur, dist, averspeed, cal), nil

}
