// boxed just prints your strings in an ascii box
package box

import (
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"
)

// Box returns a string in a box
func Box(msg string) string {
	inner := fmt.Sprintf("| %s |", msg)
	seperator := strings.Repeat("-", utf8.RuneCountInString(inner))
	return fmt.Sprintf("%s \n%s \n%s", seperator, inner, seperator)
}

// BoxMany returns many strings in a box
func BoxMany(msgs ...string) string {
	var tmp []string
	lgst := lengthiest(msgs)
	for _, msg := range msgs {
		tl := len(lgst) - len(msg)
		padding := strings.Repeat(" ", tl)
		tmp = append(tmp, fmt.Sprintf("| %s %s|\n", msg, padding))
	}
	lgstPad := lengthiest(tmp)
	seperator := strings.Repeat("-", utf8.RuneCountInString(lgstPad)-1)
	s1 := slices.Insert(tmp, 0, fmt.Sprintf("%s\n", seperator))
	s1 = append(s1, seperator)
	return strings.Join(s1[:], "")
}

// lengthiest returns an int representing the length of the longest
// string in a []string
func lengthiest(ws []string) string {
	l := 0
	var v string
	for _, w := range ws {
		ln := len(w)
		if ln > l {
			l = ln
			v = w
		}
	}
	return v
}
