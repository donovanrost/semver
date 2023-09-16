package semver

import "strconv"

type Semver struct {
	Major int
	Minor int
	Patch int
	Pre []string
	Meta []string
}

func (s Semver) String() string {
	result := ""
	result += strconv.Itoa(s.Major)
	result += "."
	result += strconv.Itoa(s.Minor)
	result += "."
	result += strconv.Itoa(s.Patch)
	
	if len(s.Pre) > 0 {
		result += "-"
		for i, p := range s.Pre {
			result += p
			if i < len(s.Pre) - 1 {
				result += "."
			}
		}
	}
	if len(s.Meta) > 0 {
		result += "+" 
		for i, p := range s.Meta {
			result += p
			if i < len(s.Meta) - 1 {
				result += "."
			}
		}
	}
	
	return result
}
type Parser interface {
	Parse(unparsed string) (Semver, error)
}

func NewFromString(unparsed string) (Semver, error) {

	parser := DefaultSemverParser{
		VParts: make([]int, 3),
		PreParts: make([]string, 0),
		MetaParts: make([]string, 0),
		
	}

	return parser.Parse(unparsed)

}
