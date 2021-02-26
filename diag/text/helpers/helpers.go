package helpers

import "regexp"

// RemoveTextMarkers removes all #tag# markers in the given string and returns new string and map of tags positions
func RemoveTextMarkers(str string) (string, map[string]int) {
	m := map[string]int{}
	runes := []rune(str)
	r := regexp.MustCompile(`#[\S\d]*#`)
	l := 0
	for _, match := range r.FindAllStringIndex(str, -1) {
		start := match[0]
		end := match[1]
		pos := start - l
		text := string(runes[start:end])
		m[text] = pos
		l += end - start
	}
	rstr := r.ReplaceAllString(str, "")
	return rstr, m
}
