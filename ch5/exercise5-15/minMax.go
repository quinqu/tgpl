package minMax

import (
	"fmt"
)

func min(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("please provide at least one argument")

	} else if len(vals) == 1 {
		fmt.Println(vals[0])
		return vals[0], nil
	}
	min := vals[0]

	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min, nil
}

func max(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("please provide at least one argument")

	} else if len(vals) == 1 {
		fmt.Println(vals[0])
		return vals[0], nil
	}
	max := vals[0]

	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max, nil
}
