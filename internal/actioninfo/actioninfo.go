package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию

	parsedOK := false

	for _, value := range dataset {
		if err := dp.Parse(value); err != nil {
			log.Println(err)
			continue
		}
		parsedOK = true
	}

	if !parsedOK {
		return
	}

	str, err := dp.ActionInfo()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Print(str)

}
