package dog

import (
	"strings"
)

func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}

func BarkAt(s string) string {
	return "Wooh Wooh, " + strings.ToTitle(s)

}
