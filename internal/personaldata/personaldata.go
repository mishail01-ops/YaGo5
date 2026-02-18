package personaldata

import "fmt"

type Personal struct {
	// TODO: добавить поля
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	// TODO: реализовать функцию
	// выводит данные структуры на экран вот в таком виде:
	// Имя: <Имя_пользователя>
	// Вес: <Вес_пользователя>
	// Рост: <Рост_пользователя>
	fmt.Println("Имя: ", p.Name)
	fmt.Println("Вес: ", p.Weight)
	fmt.Println("Рост: ", p.Height)
}
