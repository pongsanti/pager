package pager

import "strings"

const defaultColumn = "id"
const defaultDirection = "asc"

type Sort struct {
	column    string
	direction string
}

func NewSort(sort string) Sort {
	s := Sort{}

	sorts := strings.Split(sort, ",")
	if len(sorts) == 1 {
		if sorts[0] == "" {
			s.column = defaultColumn
		} else {
			s.column = sorts[0]
		}
		s.direction = defaultDirection
	} else {
		s.column = sorts[0]
		s.direction = sorts[1]
	}

	return s
}

func (s Sort) Column() string {
	return s.column
}

func (s Sort) Direction() string {
	return s.direction
}
