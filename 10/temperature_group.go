package main

import (
	"fmt"
	"sync"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. 
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

// TmGroup Splitting by groups numbers from slice
func TmGroup(tmp []float64) map[int][]float64 {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	groups := make(map[int][]float64)

	for _, val := range tmp {
		wg.Add(1)
		go func(val float64) {
			defer wg.Done()
			if val > 0 {
				for i := float64(0); i < 100; i += 10 {
					if i <= val && val < i+10 {
						mu.Lock()
						groups[int(i)] = append(groups[int(i)], val)
						mu.Unlock()
						break
					}
				}
			} else {
				for i := float64(0); i > -100; i -= 10 {
					if i >= val && val > i-10 {
						mu.Lock()
						groups[int(i)] = append(groups[int(i)], val)
						mu.Unlock()
						break
					}
				}
			}
		}(val)
	}
	wg.Wait()
	return groups
}

// TmGroup1 - another way to split numbers into groups
func TmGroup1(tmp []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, val := range tmp {
		key := int(val/10) * 10
		groups[key] = append(groups[key], val)
	}
	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	//groups := TmGroup(temperatures)
	groups := TmGroup1(temperatures)

	fmt.Printf("resul of splitting by groups %v\n", groups)
}
