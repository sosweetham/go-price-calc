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