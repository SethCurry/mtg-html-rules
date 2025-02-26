package ruleparser

// via https://github.com/Wissance/gfu/blob/master/file_utils.go
// but meh at importing full package
import (
	"io"
	"strings"
)

func ReadAllLines(reader io.Reader, omitEmpty bool) ([]string, error) {
	var err error
	var content []byte

	content, err = io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	strContent := string(content)

	rawLines := strings.Split(strContent, "\r")
	lines := make([]string, 0)

	finalLinesNumber := 0
	if !omitEmpty {
		finalLinesNumber = len(rawLines)
	}

	for _, l := range rawLines {
		l = strings.Trim(l, "\r\n")
		if omitEmpty {
			spaceTrimmedLine := strings.Trim(l, " \t")
			if len(spaceTrimmedLine) > 0 {
				lines = append(lines, l)
				finalLinesNumber++
			}
		} else {
			lines = append(lines, l)
		}
	}
	return lines[0:finalLinesNumber], err
}

func prepend[T any](x []T, y T) []T {
	var empty T
	x = append(x, empty)
	copy(x[1:], x)
	x[0] = y
	return x
}
