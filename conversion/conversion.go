package conversion

import (
	"fmt"
	"strconv"
)

func LimitFloat(initialVal float64, limit int) (float64, error) {
	if limit < 0 {
		return 0, fmt.Errorf("limit must be greater than or equal to 0")
	}
	return strconv.ParseFloat(fmt.Sprintf("%."+fmt.Sprint(limit)+"f", initialVal), 64)
}

func StringsToFloats (lines []string, limit int) ([]float64, error) {
	prices := make([]float64, len(lines))
	for _, line := range lines {
		initialVal, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		val, err := LimitFloat(initialVal, limit)
		if err != nil {
			return nil, err
		}
		prices = append(prices, val)
	}
	return prices, nil
}
