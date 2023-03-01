package enum

import (
	"fmt"
)

type Season int

var seasons = [...]string{"spring", "summer", "autumn", "winter"}

const (
	Spring Season = iota + 1
	Summer
	Autumn
	Winter
)

func (s Season) String() string {
	if s < Spring || s > Winter {
		return fmt.Sprintf("Season(%d)", int(s))
	}
	return seasons[s-1]
}

func (s Season) IsValid() bool {
	switch s {
	case Spring, Summer, Autumn, Winter:
		return true
	}
	return false
}
