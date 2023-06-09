package main

import (
	"fmt"
	"sort"
)

func main() {
	step := 10
	list := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	sort.Float64s(list)

	temperatureGroups := make(map[int][]float64)

	for _, v := range list {
		rounded := (int(v) / step) * step
		temperatureGroups[rounded] = append(temperatureGroups[rounded], v)
	}
	fmt.Println(temperatureGroups)
}
