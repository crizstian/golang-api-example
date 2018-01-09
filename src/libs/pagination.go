package libs

import (
	"strconv"
)

// GetOffsetLimit process and converts page & limit strings into int value
func GetOffsetLimit(page, limit string) (int, int) {
	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}

	l, err := strconv.Atoi(limit)
	if err != nil {
		l = 20
	}

	offset := calculateOffset(p, l)

	return offset, l
}

// calculates offset based on page and limit values
func calculateOffset(page, limit int) int {
	return (page - 1) * limit
}
