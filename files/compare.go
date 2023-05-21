package files

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"os"
)

type Line struct {
	content []byte
	number  int
}

func NewLine(content []byte, number int) Line {
	return Line{
		content: content,
		number:  number,
	}
}

type LinePair struct {
	first  Line
	second Line
}

func deepCompare(file1, file2 string) ([]LinePair, error) {
	diffs := []LinePair{}
	sf, err := os.Open(file1)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not open file %s", file1)
	}

	df, err := os.Open(file2)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not open file %s", file2)
	}

	sscan := bufio.NewScanner(sf)
	dscan := bufio.NewScanner(df)

	lineNum := 1

	for sscan.Scan() {
		line1 := NewLine(sscan.Bytes(), lineNum)
		line2 := NewLine(dscan.Bytes(), lineNum)
		lineNum++
		if !bytes.Equal(sscan.Bytes(), dscan.Bytes()) {
			diffs = append(diffs, LinePair{
				line1,
				line2,
			})
		}
	}

	for dscan.Scan() {
		line1 := NewLine([]byte{}, -1)
		line2 := NewLine(dscan.Bytes(), lineNum)
		lineNum++
		diffs = append(diffs, LinePair{
			line1,
			line2,
		})
	}
	return diffs, nil
}
