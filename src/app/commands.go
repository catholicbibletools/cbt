package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	. "../bibles"
)

type CommandResult struct {
	Info      string
	Prelude   string
	Content   *list.List
	Corollary string
	Error     string
}

func NewCommandResult() *CommandResult {
	return &CommandResult{"", "", list.New(), "", ""}
}

func (me *CommandResult) Add(content string) {
	me.Content.PushBack(content)
}

func ProcessCommandResult(result *CommandResult) int {
	if result.Error != "" {
		fmt.Println("ERROR: " + result.Error)
		return 1
	}
	if result.Info != "" {
		fmt.Println(result.Info)
	}
	if result.Prelude != "" {
		fmt.Println(result.Prelude)
	}
	for e := result.Content.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	if result.Corollary != "" {
		fmt.Println(result.Corollary)
	}
	return 0
}

func CommandVersion() *CommandResult {
	result := NewCommandResult()
	result.Add(TITLE + " " + VERSION)
	return result
}

func CommandHelp(cmd string) *CommandResult {
	result := NewCommandResult()
	switch cmd {
	case "v":
		result.Add("desription:")
		result.Add("  shows the current version")
		result.Add("usage:")
		result.Add("  cbt v")
		return result
	case "l":
		result.Add("desription:")
		result.Add("  lists the available bibles")
		result.Add("usage:")
		result.Add("  cbt l")
		return result
	case "c":
		result.Add("desription:")
		result.Add("  displays current configuration settings or sets specific setting to the current configuration")
		result.Add("usage:")
		result.Add("  cbt c [<setting> [<value>]]")
		result.Add("examples:")
		result.Add("  cbt c")
		result.Add("    displays all the current configuration settings")
		result.Add("  cbt c current-bible")
		result.Add("    displays the current bible in the current configuration settings")
		result.Add("  cbt c current-bible en.ccb")
		result.Add("    sets the current bible to 'en.ccb' in the current configuration settings")
		return result
	case "r":
		result.Add("desription:")
		result.Add("  reads the bible")
		result.Add("usage:")
		result.Add("  cbt r <what>")
		result.Add("examples:")
		result.Add("  cbt r 7 gn")
		result.Add("    reads Genesis chapter 7")
		result.Add("  cbt r 1 jn 1,1")
		result.Add("    reads 1 John chapter 1 versicle 1")
		result.Add("  cbt r 2 co 3,1-5")
		result.Add("    reads 2 Corinthians chapter 2 versicles 1 to 5")
		result.Add("  cbt r ap 1-3")
		result.Add("    reads Apokalypsis chapters 1 to 3")
		result.Add("  cbt r mt 10,1-12,2")
		result.Add("    reads Matthew from versicle 10,1 to 12,2")
		return result
	case "s":
		result.Add("desription:")
		result.Add("  searches the bible")
		result.Add("usage:")
		result.Add("  cbt s '<what>' [-'<excluding>'] [<where>]")
		result.Add("examples:")
		result.Add("  cbt s 'love is'")
		result.Add("    searches all the Bible for 'love is'")
		result.Add("  cbt s 'love is' o")
		result.Add("    searches the Old Testament for 'love is'")
		result.Add("  cbt s 'love is' n")
		result.Add("    searches the New Testament for 'love is'")
		result.Add("  cbt s 'love is' ap")
		result.Add("    searches the Apokalypsis for 'love is'")
		result.Add("  cbt s 'love is' ap 1")
		result.Add("    searches the Apokalypsis chapter 1 for 'love is'")
		result.Add("  cbt s 'love is' 1 jn 1")
		result.Add("    searches 1 John chapter 1 for 'love is'")
		result.Add("  cbt s 'the first' -'fire' ap")
		result.Add("    searches the Apokalypsis for 'the first' excluding verses where 'over them' appears")
		return result
	case "w":
		result.Add("desription:")
		result.Add("  starts the web api")
		result.Add("usage:")
		result.Add("  cbt w <port>")
		return result
	}
	result.Add("desription:")
	result.Add("  this information")
	result.Add("usage:")
	result.Add("  cbt h <command>")
	result.Add("commands:")
	result.Add("  v - version")
	result.Add("  h - help")
	result.Add("  l - list")
	result.Add("  c - config")
	result.Add("  r - read")
	result.Add("  s - search")
	result.Add("  w - web")
	return result
}

func CommandList() *CommandResult {
	result := NewCommandResult()
	result.Info = "-- available bibles:"
	for x := range Bibles {
		b := Bibles[x]
		result.Add(b.GetCode() + " - " + b.GetName())
	}
	return result
}

func CommandConfig(key string, value string) *CommandResult {
	result := NewCommandResult()
	if key == "" {
		// display settings
		result.Add("current-language: '" + currentConfiguration.GetCurrentLanguage() + "'")
		result.Add("current-bible: '" + currentConfiguration.GetCurrentBible() + "'")
		result.Add("output-format: '" + currentConfiguration.GetOutputFormat() + "'")
		result.Add("outer-separator: '" + currentConfiguration.GetOuterSeparator() + "'")
		result.Add("major-separator: '" + currentConfiguration.GetMajorSeparator() + "'")
		result.Add("minor-separator: '" + currentConfiguration.GetMinorSeparator() + "'")
	} else {
		if value != "" {
			// set setting
			switch key {
			case "current-language":
				if IsValidLanguage(value) {
					currentConfiguration.SetCurrentLanguage(value)
					currentAbbreviations = GetAbbreviations(currentConfiguration.GetCurrentLanguage())
				} else {
					result.Error = "unknown language: '" + value + "'"
				}
			case "current-bible":
				if IsValidBible(value) {
					currentConfiguration.SetCurrentBible(value)
					currentBible = GetBible(currentConfiguration.GetCurrentBible())
					currentBible.LoadBooks()
				} else {
					result.Error = "unknown bible: '" + value + "'"
				}
			case "output-format":
				if IsValidFormat(value) {
					currentConfiguration.SetOutputFormat(value)
					LoadTemplates(currentConfiguration.GetOutputFormat())
				} else {
					result.Error = "unknown format: '" + value + "'"
				}
			case "outer-separator":
				currentConfiguration.SetOuterSeparator(value)
			case "major-separator":
				currentConfiguration.SetMajorSeparator(value)
			case "minor-separator":
				currentConfiguration.SetMinorSeparator(value)
			}
		}
		// get setting
		switch key {
		case "current-language":
			result.Add("current-language: " + currentConfiguration.GetCurrentLanguage())
		case "current-bible":
			result.Add("current-bible: " + currentConfiguration.GetCurrentBible())
		case "output-format":
			result.Add("output-format: " + currentConfiguration.GetOutputFormat())
		case "outer-separator":
			result.Add("outer-separator: " + currentConfiguration.GetOuterSeparator())
		case "major-separator":
			result.Add("major-separator: " + currentConfiguration.GetMajorSeparator())
		case "minor-separator":
			result.Add("minor-separator: " + currentConfiguration.GetMinorSeparator())
		default:
			result.Error = "unknown setting: '" + key + "'"
		}
	}
	return result
}

func CommandRead(input string) *CommandResult {
	psg := NewPassage(input, currentAbbreviations)
	result := NewCommandResult()
	result.Info = "-- read input '" + input + "' interpreted as: "
	if psg.BookCode == "" {
		result.Error = "can not understand target '" + input + "'"
	} else {
		bk := currentBible.GetBook(psg.BookCode)
		var at string = bk.Name
		if psg.BeginChapter == 0 {
			// full-book read
			if bk.BeginChapter == bk.EndChapter {
				at += " " + strconv.Itoa(bk.BeginChapter)
				result.Info += at + " (read type: full-book single-chapter)"
				cn := strconv.Itoa(bk.BeginChapter)
				bc := bk.GetChapter(bk.BeginChapter)
				if bc == nil {
					result.Error = "there is no chapter " + cn
					return result
				}
				if ol := ReadOpenFormat(currentConfiguration, bk.Name); ol != "" {
					result.Add(ol)
				}
				vs := bc.GetVersicleRange(bc.BeginVersicle, bc.EndVersicle)
				for x := range vs {
					result.Add(ReadVersiclesFormat(currentConfiguration, bk.Name, cn, vs[x].Number, vs[x].Text))
					//result.Add(vs[x].Number + " " + vs[x].Text)
				}
				if cl := ReadCloseFormat(currentConfiguration, bk.Name); cl != "" {
					result.Add(cl)
				}
			} else {
				at += " " + strconv.Itoa(bk.BeginChapter) + currentConfiguration.GetMinorSeparator() + strconv.Itoa(bk.EndChapter)
				result.Info += at + " (read type: full-book multi-chapter)"
				if ol := ReadOpenFormat(currentConfiguration, bk.Name); ol != "" {
					result.Add(ol)
				}
				for i := bk.BeginChapter; i <= bk.EndChapter; i++ {
					//result.Add(strconv.Itoa(i))
					cn := strconv.Itoa(i)
					bc := bk.GetChapter(i)
					if bc == nil {
						result.Error = "there is no chapter " + cn
						return result
					}
					vs := bc.GetVersicleRange(bc.BeginVersicle, bc.EndVersicle)
					for x := range vs {
						result.Add(ReadVersiclesFormat(currentConfiguration, bk.Name, cn, vs[x].Number, vs[x].Text))
						//result.Add(vs[x].Number + " " + vs[x].Text)
					}
				}
				if cl := ReadCloseFormat(currentConfiguration, bk.Name); cl != "" {
					result.Add(cl)
				}
			}
		} else {
			// chapter read
			bc := bk.GetChapter(psg.BeginChapter)
			if bc == nil {
				result.Error = "there is no chapter " + strconv.Itoa(psg.BeginChapter)
				return result
			}
			if ol := ReadOpenFormat(currentConfiguration, bk.Name); ol != "" {
				result.Add(ol)
			}
			if psg.BeginChapter == psg.EndChapter {
				// single-chapter read
				cn := strconv.Itoa(psg.BeginChapter)
				if psg.BeginVersicle == "" && psg.EndVersicle == "" {
					at += " " + cn
					result.Info += at + " (read type: single-chapter full-chapter)"
					vs := bc.GetVersicleRange(bc.BeginVersicle, bc.EndVersicle)
					for x := range vs {
						result.Add(ReadVersiclesFormat(currentConfiguration, bk.Name, cn, vs[x].Number, vs[x].Text))
						//result.Add(vs[x].Number + " " + vs[x].Text)
					}
				} else {
					if psg.EndVersicle == "" {
						at += " " + cn + currentConfiguration.GetMajorSeparator() + psg.BeginVersicle
						result.Info += at + " (read type: single-chapter single-versicle)"
						result.Add(ReadVersiclesFormat(currentConfiguration, bk.Name, cn, psg.BeginVersicle, bc.GetVersicleText(psg.BeginVersicle)))
						//result.Add(psg.BeginVersicle + " " + bc.GetVersicleText(psg.BeginVersicle))
					} else {
						at += " " + cn + currentConfiguration.GetMajorSeparator() + psg.BeginVersicle + currentConfiguration.GetMinorSeparator() + psg.EndVersicle
						result.Info += at + " (read type: single-chapter multi-versicle)"
						vs := bc.GetVersicleRange(psg.BeginVersicle, psg.EndVersicle)
						for x := range vs {
							result.Add(ReadVersiclesFormat(currentConfiguration, bk.Name, cn, vs[x].Number, vs[x].Text))
							//result.Add(vs[x].Number + " " + vs[x].Text)
						}
					}
				}
			} else {
				// multi-chapter read
				at += " " + strconv.Itoa(psg.BeginChapter)
				if psg.BeginVersicle != "" {
					at += currentConfiguration.GetMajorSeparator() + psg.BeginVersicle
				}
				at += currentConfiguration.GetMinorSeparator() + strconv.Itoa(psg.EndChapter)
				if psg.EndVersicle != "" {
					at += currentConfiguration.GetMajorSeparator() + psg.EndVersicle
				}
				result.Info += at + " (read type: multi-chapter)"
				for i := psg.BeginChapter; i <= psg.EndChapter; i++ {
					cn := strconv.Itoa(i)
					bc := bk.GetChapter(i)
					if bc == nil {
						result.Error = "there is no chapter " + strconv.Itoa(i)
						break
					}
					//result.Add(strconv.Itoa(i))
					var from string = bc.BeginVersicle
					var to string = bc.EndVersicle
					if i == psg.BeginChapter && psg.BeginVersicle != "" {
						from = psg.BeginVersicle
					} else if i == psg.EndChapter && psg.EndVersicle != "" {
						to = psg.EndVersicle
					}
					vs := bc.GetVersicleRange(from, to)
					for x := range vs {
						result.Add(ReadVersiclesFormat(currentConfiguration, bk.Name, cn, vs[x].Number, vs[x].Text))
						//result.Add(vs[x].Number + " " + vs[x].Text)
					}
				}
			}
			if cl := ReadCloseFormat(currentConfiguration, bk.Name); cl != "" {
				result.Add(cl)
			}
		}
	}
	return result
}

func CommandSearch(text string, exclude string, input string) *CommandResult {
	result := NewCommandResult()
	if exclude != "" {
		result.Info = "-- searching for '" + text + "' excluding '" + exclude + "' at:"
	} else {
		result.Info = "-- searching for '" + text + "' at:"
	}
	switch input {
	case "":
		result.Info += " Bible (search type: multi-testament)"
		if ol := TestamentsSearchOpenFormat(currentConfiguration); ol != "" {
			result.Add(ol)
		}
		hits := 0
		for x := range currentBible.Codes {
			bk := currentBible.GetBook(currentBible.Codes[x])
			for i := bk.BeginChapter; i <= bk.EndChapter; i++ {
				cn := strconv.Itoa(i)
				bc := bk.GetChapter(i)
				vs := bc.GetVersicleRange(bc.BeginVersicle, bc.EndVersicle)
				for x := range vs {
					if strings.Contains(vs[x].Text, text) {
						if exclude != "" && strings.Contains(vs[x].Text, exclude) {
							continue
						}
						hits++
						hit := bk.Name + " " + cn + currentConfiguration.GetMajorSeparator() + vs[x].Number
						result.Add(TestamentsSearchVersiclesFormat(currentConfiguration, bk.Name, cn, vs[x].Number, vs[x].Text))
						//result.Add(hit + " " + vs[x].Text)
						result.Corollary += currentConfiguration.GetOuterSeparator() + " " + hit
					}
				}
			}
		}
		if cl := TestamentsSearchCloseFormat(currentConfiguration); cl != "" {
			result.Add(cl)
		}
		result.Prelude = "Found: " + strconv.Itoa(hits) + " times"
		if hits > 0 {
			result.Corollary = "Hits:" + strings.TrimLeft(result.Corollary, currentConfiguration.GetOuterSeparator())
		}
	case "o", "v":
		result.Info += " Old Testament (search type: single-testament)"
		if ol := TestamentsSearchOpenFormat(currentConfiguration); ol != "" {
			result.Add(ol)
		}
		hits := 0
		for x := range currentBible.Codes {
			c := currentBible.Codes[x]
			if strings.Index(c, "V") == 0 {
				bk := currentBible.GetBook(c)
				for i := bk.BeginChapter; i <= bk.EndChapter; i++ {
					cn := strconv.Itoa(i)
					bc := bk.GetChapter(i)
					vs := bc.GetVersicleRange(bc.BeginVersicle, bc.EndVersicle)
					for x := range vs {
						if strings.Contains(vs[x].Text, text) {
							if exclude != "" && strings.Contains(vs[x].Text, exclude) {
								continue
							}
							hits++
							hit := bk.Name + " " + cn + currentConfiguration.GetMajorSeparator() + vs[x].Number
							result.Add(TestamentsSearchVersiclesFormat(currentConfiguration, bk.Name, cn, vs[x].Number, vs[x].Text))
							//result.Add(hit + " " + vs[x].Text)
							result.Corollary += currentConfiguration.GetOuterSeparator() + " " + hit
						}
					}
				}
			}
		}
		if cl := TestamentsSearchCloseFormat(currentConfiguration); cl != "" {
			result.Add(cl)
		}
		result.Prelude = "Found: " + strconv.Itoa(hits) + " times"
		if hits > 0 {
			result.Corollary = "Hits:" + strings.TrimLeft(result.Corollary, currentConfiguration.GetOuterSeparator())
		}
	case "n":
		result.Info += " New Testament (search type: single-testament)"
		if ol := TestamentsSearchOpenFormat(currentConfiguration); ol != "" {
			result.Add(ol)
		}
		hits := 0
		for x := range currentBible.Codes {
			c := currentBible.Codes[x]
			if strings.Index(c, "N") == 0 {
				bk := currentBible.GetBook(c)
				for i := bk.BeginChapter; i <= bk.EndChapter; i++ {
					cn := strconv.Itoa(i)
					bc := bk.GetChapter(i)
					vs := bc.GetVersicleRange(bc.BeginVersicle, bc.EndVersicle)
					for x := range vs {
						if strings.Contains(vs[x].Text, text) {
							if exclude != "" && strings.Contains(vs[x].Text, exclude) {
								continue
							}
							hits++
							hit := bk.Name + " " + cn + currentConfiguration.GetMajorSeparator() + vs[x].Number
							result.Add(TestamentsSearchVersiclesFormat(currentConfiguration, bk.Name, cn, vs[x].Number, vs[x].Text))
							//result.Add(hit + " " + vs[x].Text)
							result.Corollary += currentConfiguration.GetOuterSeparator() + " " + hit
						}
					}
				}
			}
		}
		if cl := TestamentsSearchCloseFormat(currentConfiguration); cl != "" {
			result.Add(cl)
		}
		result.Prelude = "Found: " + strconv.Itoa(hits) + " times"
		if hits > 0 {
			result.Corollary = "Hits:" + strings.TrimLeft(result.Corollary, currentConfiguration.GetOuterSeparator())
		}
	default:
		psg := NewPassage(input, currentAbbreviations)
		if psg.BookCode == "" {
			result.Error = "can not understand target '" + input + "'"
		} else {
			bk := currentBible.GetBook(psg.BookCode)
			result.Info += " " + bk.Name
			if psg.BeginChapter == 0 {
				// book search
				result.Info += " " + strconv.Itoa(bk.BeginChapter) + currentConfiguration.GetMinorSeparator() + strconv.Itoa(bk.EndChapter) + " (search type: single-book)"
				if ol := BookSearchOpenFormat(currentConfiguration, bk.Name); ol != "" {
					result.Add(ol)
				}
				hits := 0
				for i := bk.BeginChapter; i <= bk.EndChapter; i++ {
					cn := strconv.Itoa(i)
					bc := bk.GetChapter(i)
					vs := bc.GetVersicleRange(bc.BeginVersicle, bc.EndVersicle)
					for x := range vs {
						if strings.Contains(vs[x].Text, text) {
							if exclude != "" && strings.Contains(vs[x].Text, exclude) {
								continue
							}
							hits++
							hit := bk.Name + " " + cn + currentConfiguration.GetMajorSeparator() + vs[x].Number
							result.Add(BookSearchVersiclesFormat(currentConfiguration, bk.Name, cn, vs[x].Number, vs[x].Text))
							//result.Add( /*bk.Name + " " + */ cn + currentConfiguration.GetMajorSeparator() + vs[x].Number + " " + vs[x].Text)
							result.Corollary += currentConfiguration.GetOuterSeparator() + " " + hit
						}
					}
				}
				if cl := BookSearchCloseFormat(currentConfiguration, bk.Name); cl != "" {
					result.Add(cl)
				}
				result.Prelude = "Found: " + strconv.Itoa(hits) + " times"
				if hits > 0 {
					result.Corollary = "Hits:" + strings.TrimLeft(result.Corollary, currentConfiguration.GetOuterSeparator())
				}
			} else {
				// chapter search
				cn := strconv.Itoa(psg.BeginChapter)
				result.Info += " " + cn + " (search type: single-chapter)"
				bc := bk.GetChapter(psg.BeginChapter)
				if bc == nil {
					result.Error = "there is no chapter " + cn
					break
				}
				if ol := ChapterSearchOpenFormat(currentConfiguration, bk.Name, cn); ol != "" {
					result.Add(ol)
				}
				vs := bc.GetVersicleRange(bc.BeginVersicle, bc.EndVersicle)
				hits := 0
				for x := range vs {
					if strings.Contains(vs[x].Text, text) {
						if exclude != "" && strings.Contains(vs[x].Text, exclude) {
							continue
						}
						hits++
						hit := bk.Name + " " + cn + currentConfiguration.GetMajorSeparator() + vs[x].Number
						result.Add(ChapterSearchVersiclesFormat(currentConfiguration, bk.Name, cn, vs[x].Number, vs[x].Text))
						//result.Add( /*bk.Name + " " + cn + MajorSeparator + */ vs[x].Number + " " + vs[x].Text)
						result.Corollary += currentConfiguration.GetOuterSeparator() + " " + hit
					}
				}
				if cl := ChapterSearchCloseFormat(currentConfiguration, bk.Name, cn); cl != "" {
					result.Add(cl)
				}
				result.Prelude = "Found: " + strconv.Itoa(hits) + " times"
				if hits > 0 {
					result.Corollary = "Hits:" + strings.TrimLeft(result.Corollary, currentConfiguration.GetOuterSeparator())
				}
			}
		}
	}
	return result
}

func ProcessCommandResultForWeb(result *CommandResult, w http.ResponseWriter, literalContent bool) {
	reply := "{"
	if result.Error != "" {
		reply += "\"error\":\"" + result.Error + "\""
	} else {
		if result.Info != "" {
			reply += "\"info\":\"" + result.Info + "\","
		}
		if literalContent {
			reply += "\"content\":"
			for e := result.Content.Front(); e != nil; e = e.Next() {
				reply += fmt.Sprintf("%v", e.Value) + "\n"
			}
		} else {
			reply += "\"content\":["
			if result.Prelude != "" {
				reply += "\"" + result.Prelude + "\",\n"
			}
			first := true
			for e := result.Content.Front(); e != nil; e = e.Next() {
				if first {
					first = false
				} else {
					reply += ",\n"
				}
				reply += "\"" + fmt.Sprintf("%v", e.Value) + "\""
			}
			if result.Corollary != "" {
				reply += "\"" + result.Corollary + "\"\n"
			}
			reply += "]"
		}
	}
	w.Write([]byte(reply + "}\n"))
}

func globalHandler(w http.ResponseWriter, r *http.Request) {
	url := strings.TrimRight(r.URL.Path[1:], "/")
	switch url {
	case "version":
		ProcessCommandResultForWeb(CommandVersion(), w, false)
	case "list":
		ProcessCommandResultForWeb(CommandList(), w, false)
	case "config":
		body, err1 := ioutil.ReadAll(r.Body)
		if err1 != nil {
			//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			ProcessCommandResultForWeb(CommandConfig("", ""), w, false)
			return
		}
		t := struct {
			Setting string `json:"setting"`
			Value   string `json:"value"`
		}{}
		err2 := json.Unmarshal(body, &t)
		if err2 != nil {
			//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			ProcessCommandResultForWeb(CommandConfig("", ""), w, false)
			return
		}
		ProcessCommandResultForWeb(CommandConfig(t.Setting, t.Value), w, false)
	case "read":
		body, err1 := ioutil.ReadAll(r.Body)
		if err1 != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		t := struct {
			What string `json:"what"`
		}{}
		err2 := json.Unmarshal(body, &t)
		if err2 != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		ProcessCommandResultForWeb(CommandRead(t.What), w, currentConfiguration.GetOutputFormat() == "json")
	case "search":
		body, err1 := ioutil.ReadAll(r.Body)
		if err1 != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		t := struct {
			What      string `json:"what"`
			Excluding string `json:"excluding"`
			Where     string `json:"where"`
		}{}
		err2 := json.Unmarshal(body, &t)
		if err2 != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		ProcessCommandResultForWeb(CommandSearch(t.What, t.Excluding, t.Where), w, currentConfiguration.GetOutputFormat() == "json")
	default:
		http.Redirect(w, r, "/client/", http.StatusSeeOther)
	}
}

func CommandWeb(port string) *CommandResult {
	result := NewCommandResult()
	result.Info = "-- started web api at port '" + port
	mux := http.NewServeMux()
	mux.HandleFunc("/", globalHandler)
	fs := http.FileServer(http.Dir("client"))
	mux.Handle("/client/", http.StripPrefix("/client/", fs))
	http.ListenAndServe(":"+port, mux)
	return result
}
