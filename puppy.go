package puppy

import (
	"github.com/hanazarraa/dog"
)

func Bark() string {
	return "woof"
}
func Barks() string {
	return "woof woof woof"
}
func BigBark() string {
	return dog.WhenGrowUp(Bark())
}
func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}
