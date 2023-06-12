package main

import (
	"fmt"
	"sort"
)

func main() {
	step := 10                                                           // Шаг для группировки температур
	list := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // Список температур

	sort.Float64s(list) // Сортировка списка температур по возрастанию

	temperatureGroups := make(map[int][]float64) // Создание словаря для группировки температур

	for _, v := range list {
		rounded := (int(v) / step) * step                                  // Округление температуры до ближайшего кратного шага
		temperatureGroups[rounded] = append(temperatureGroups[rounded], v) // Добавление температуры в соответствующую группу
	}

	fmt.Println(temperatureGroups) // Вывод группированных температур
}
