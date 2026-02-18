package daysteps

import (
	"errors"
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

	var ers error
	var step int
	var dur time.Duration
	ers = nil
	spl := strings.Split(datastring, ",")
	if len(spl) != 2 {
		ers = errors.New("Данные неправильного формата, должно 678,0h50m")
		return ers
	}

	step, ers = strconv.Atoi(spl[0])

	if ers != nil {
		return ers
	}

	if step <= 0 {
		ers = errors.New("Количество шагов должно быть больше 0")
		return ers
	}

	dur, ers = time.ParseDuration(spl[1])

	if ers != nil {
		return ers
	}

	if dur <= 0 {
		ers = errors.New("Длительность должна быть больше 0")
		return ers
	}

	ds.Steps = step
	ds.Duration = dur
	return ers
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}
	//Количество шагов: 792.
	//Дистанция составила 0.51 км.
	//Вы сожгли 221.33 ккал.
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cal), nil
}
