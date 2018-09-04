package db

import "fmt"

type UpdateSection struct {
	Name  string
	Value string
}

func (u UpdateSection) Merge() string {
	return fmt.Sprintf("%s='%s'", u.Name, u.Value)
}
