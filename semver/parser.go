package semver

import (
	"errors"
	"strconv"
	"unicode"
)

type DefaultSemverParser struct {
	VParts []int
	PreParts []string
	MetaParts []string
	errPos int
	input []rune
}

func (d *DefaultSemverParser) Parse(unparsed string ) (Semver, error ) {
	d.input = []rune(unparsed)
	err := d.stateMajor()
	if err != nil {
		return Semver{}, err
	}

	return Semver{
		Major: d.VParts[0],
		Minor: d.VParts[1],
		Patch: d.VParts[2],
		Pre: d.PreParts,
		Meta: d.MetaParts,

		}, nil
}

func (d *DefaultSemverParser) stateMajor() error {
	pos := 0

	for pos < len(d.input) && isDigit(d.input[pos]) {
		pos++
	}
	if pos == 0 {
		d.errPos = pos
		return errors.New("empty input")
	}
	if d.input[0] == '0' && pos > 1 {
		d.errPos = pos
		return errors.New("leading 0")
	}

	substr := string(d.input[0 : pos ])
	major, _ := strconv.Atoi(substr) 

	d.VParts[0] = major

	if d.input[pos] == '.' {
		return d.stateMinor(pos + 1)
	}


	d.errPos = pos
	return errors.New("incomplete version string")

}

func (d *DefaultSemverParser) stateMinor (index  int) error {
	pos := index

	for pos < len(d.input) && isDigit(d.input[pos]) {
		pos++
	}
	if pos == index { 
		d.errPos = index
		return errors.New("empty input")
	}
	if d.input[0] == '0' && pos > 1 {
		d.errPos = pos
		return errors.New("leading 0")
	}

	substr := string(d.input[index : pos])
	minor, _ := strconv.Atoi(substr)

	d.VParts[1] = minor

	if d.input[pos] == '.' {
		return d.statePatch(pos + 1)
	}

	d.errPos = pos
	return errors.New("incomplete version string")
}


func (d *DefaultSemverParser) statePatch(index int ) error {
	pos := index
	for pos < len(d.input) && isDigit(d.input[pos]) {
		pos++
	}

	if pos == index { 
		d.errPos = index
		return errors.New("empty input")
	}
	if d.input[0] == '0' && pos > 1 {
		d.errPos = pos
		return errors.New("leading 0")
	}

	substr := string(d.input[index : pos])
	patch, _ := strconv.Atoi(substr)

	d.VParts[2] = patch

	if pos == len(d.input) {
		return nil // the version string is fully parsed
	}

	if d.input[pos] == '+' {
		return d.stateMeta(pos + 1)
	}
	if d.input[pos] == '-' {
		return d.statePre(pos + 1)
	}

	d.errPos = pos
	return errors.New("invalid string")
}

func (d *DefaultSemverParser) stateMeta(index int) error {
	pos := index
	for pos < len(d.input) && (isAlphanumeric(d.input[pos]) || isDash(d.input[pos])) {
		pos++
	}
	if pos == index { 
		d.errPos = index
		return errors.New("empty input")
	}

	part := string(d.input[index : pos])
	d.MetaParts = append(d.MetaParts, part)

	if pos == len(d.input) {
		return nil
	}

	if d.input[pos] == '.' {
		return d.stateMeta(pos + 1)
	}

	d.errPos = pos
	return errors.New("invalid meta")
}

func (d *DefaultSemverParser) statePre(index int) error {
	pos := index

	for pos < len(d.input) && (isAlphanumeric(d.input[pos]) || isDash(d.input[pos])) {
		pos++
	}
	if pos == index { 
		d.errPos = index
		return errors.New("empty input")
	}

	part := string(d.input[index : pos])
	d.PreParts = append(d.PreParts, part)

	if pos == len(d.input) {
		return nil
	}

	if d.input[pos] == '.' {
		return d.statePre(pos + 1)
	}

	if d.input[pos] == '+' {
		return d.stateMeta(pos + 1)
	}
	d.errPos = pos
	return errors.New("invalid prerelease")

}


func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}
func isAlphanumeric(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}

func isDash(r rune) bool {
	return r == '-'
}