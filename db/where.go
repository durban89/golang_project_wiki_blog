package db

import "fmt"

type Where struct {
	Name     string
	Value    string
	Operator string
}

func (w Where) Merge() string {
	if w.Operator == "" {
		return fmt.Sprintf("%s = %s", w.Name, w.Value)
	}
	return fmt.Sprintf("%s %s %s", w.Name, w.Operator, w.Value)
}
