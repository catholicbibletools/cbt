package bibles

import (
	"bytes"
	"encoding/json"
)

type Chapter struct {
	Number        int
	Versicles     []Versicle
	BeginVersicle string
	EndVersicle   string
}

func NewChapter(number int, content string) Chapter {
	cp := Chapter{}
	cp.Number = number
	d := json.NewDecoder(bytes.NewReader([]byte(content)))
	t, err0 := d.Token()
	if err0 != nil {
		panic(err0)
	}
	if t != json.Delim('{') {
		panic("expected start of object")
	}
	cp.Versicles = []Versicle{}
	for {
		t, err1 := d.Token()
		if err1 != nil {
			panic(err1)
		}
		if t == json.Delim('}') {
			break
		}
		v, err2 := d.Token()
		if err2 != nil {
			panic(err2)
		}
		if cp.BeginVersicle == "" {
			cp.BeginVersicle = t.(string)
		}
		cp.EndVersicle = t.(string)
		cp.Versicles = append(cp.Versicles, Versicle{
			Number: t.(string),
			Text:   v.(string),
		})
	}
	return cp
}

func (me *Chapter) GetVersicleText(number string) string {
	for n := range me.Versicles {
		var v = me.Versicles[n]
		if v.Number == number {
			return v.Text
		}
	}
	return ""
}

func (me *Chapter) GetVersicleRange(from string, to string) []Versicle {
	result := []Versicle{}
	copy := false
	for n := range me.Versicles {
		v := me.Versicles[n]
		if v.Number == from {
			copy = true
		}
		if copy == true {
			result = append(result, v)
		}
		if v.Number == to {
			break
		}
	}
	return result
}
