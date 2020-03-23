package bibles

type Book struct {
	Code         string
	Name         string
	Chapters     []Chapter
	BeginChapter int
	EndChapter   int
}

func NewBook(code string, name string) *Book {
	return &Book{code, name, []Chapter{}, 0, 0}
}

func (me *Book) AddChapter(number int, content string) {
	if me.BeginChapter == 0 {
		me.BeginChapter = number
	}
	me.EndChapter = number
	me.Chapters = append(me.Chapters, NewChapter(number, content))
}

func (me *Book) GetChapter(number int) *Chapter {
	for n := range me.Chapters {
		var c = me.Chapters[n]
		if c.Number == number {
			return &c
		}
	}
	return nil
}
