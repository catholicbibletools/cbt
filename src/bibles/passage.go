package bibles

import (
	"strconv"
	"strings"
)

type Passage struct {
	BookCode      string
	BeginChapter  int
	BeginVersicle string
	EndChapter    int
	EndVersicle   string
}

func NewPassage(passage string, abbreviations *Abbreviations) *Passage {
	result := Passage{}
	parts := strings.Split(passage, " ")
	bpart := strings.ToLower(parts[0])
	if bpart == "1" || bpart == "2" || bpart == "3" || bpart == "4" || bpart == "5" || bpart == "i" || bpart == "ii" || bpart == "iii" || bpart == "iv" || bpart == "v" {
		bpart += strings.ToLower(parts[1])
		parts = parts[2:]
	} else {
		parts = parts[1:]
	}
	result.BookCode = GetBookCodeByAbbreviation(abbreviations, bpart)
	if result.BookCode != "" && len(parts) > 0 {
		cp := parts[0]
		z := strings.Index(cp, ",")
		x := strings.Index(cp, "-")
		if z > -1 {
			// versicle centric
			if x > -1 {
				// has range
				cparts := strings.Split(cp, "-")
				r1vparts := strings.Split(cparts[0], ",")
				i1, err1 := strconv.ParseInt(r1vparts[0], 10, 32)
				if err1 != nil {
					panic(err1)
				}
				result.BeginChapter = int(i1)
				if len(r1vparts) > 1 {
					result.BeginVersicle = r1vparts[1]
				} else {
					result.BeginVersicle = "1"
				}
				r2vparts := strings.Split(cparts[1], ",")
				if len(r2vparts) > 1 {
					// multi chapter
					i2, err2 := strconv.ParseInt(r2vparts[0], 10, 32)
					if err2 != nil {
						panic(err2)
					}
					result.EndChapter = int(i2)
					result.EndVersicle = r2vparts[1]
				} else {
					// single chapter
					result.EndChapter = int(i1)
					result.EndVersicle = r2vparts[0]
				}
			} else {
				// no range
				vparts := strings.Split(cp, ",")
				i, err := strconv.ParseInt(vparts[0], 10, 32)
				if err != nil {
					panic(err)
				}
				result.BeginChapter = int(i)
				result.BeginVersicle = vparts[1]
				result.EndChapter = int(i)
			}
		} else {
			// chapter centric
			if x > -1 {
				// has range
				cparts := strings.Split(cp, "-")
				i1, err1 := strconv.ParseInt(cparts[0], 10, 32)
				if err1 != nil {
					panic(err1)
				}
				result.BeginChapter = int(i1)
				i2, err2 := strconv.ParseInt(cparts[1], 10, 32)
				if err2 != nil {
					panic(err2)
				}
				result.EndChapter = int(i2)
			} else {
				// no range
				i, err := strconv.ParseInt(cp, 10, 32)
				if err != nil {
					panic(err)
				}
				result.BeginChapter = int(i)
				result.EndChapter = int(i)
			}
		}
	}
	return &result
}
