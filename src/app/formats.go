package main

import (
	"bytes"
	"html/template"
)

func IsValidFormat(format string) bool {
	switch format {
	case "csv", "html", "json", "md", "txt", "xml":
		return true
	}
	return false
}

type Info struct {
	OuterSeparator string
	MajorSeparator string
	MinorSeparator string
	BookName       string
	ChapterNumber  string
	VersicleNumber string
	VersicleText   template.HTML
}

type FormatTemplates struct {
	templates     *template.Template
	open_tmpl     string
	versicle_tmpl string
	close_tmpl    string
}

var testamentsSearchTemplates *FormatTemplates
var bookSearchTemplates *FormatTemplates
var chapterSearchTemplates *FormatTemplates
var readTemplates *FormatTemplates

func loadFormatTemplates(outputFormat string, actionName string) *FormatTemplates {
	result := FormatTemplates{
		template.New(actionName),
		actionName + "_open.tmpl",
		actionName + "_versicle.tmpl",
		actionName + "_close.tmpl",
	}
	var err error
	result.templates, err = result.templates.ParseFiles("./templates/"+outputFormat+"/"+result.open_tmpl, "./templates/"+outputFormat+"/"+result.versicle_tmpl, "./templates/"+outputFormat+"/"+result.close_tmpl)
	if err != nil {
		panic(err)
	}
	return &result
}

func LoadTemplates(outputFormat string) {
	testamentsSearchTemplates = loadFormatTemplates(outputFormat, "testaments_search")
	bookSearchTemplates = loadFormatTemplates(outputFormat, "book_search")
	chapterSearchTemplates = loadFormatTemplates(outputFormat, "chapter_search")
	readTemplates = loadFormatTemplates(outputFormat, "read")
}

func formatOpen(t *FormatTemplates, cfg Configuration, bookName string, chapterNumber string, versicleNumber string, versicleText string) string {
	info := Info{cfg.GetOuterSeparator(), cfg.GetMajorSeparator(), cfg.GetMinorSeparator(), bookName, chapterNumber, versicleNumber, template.HTML(versicleText)}
	var tpl bytes.Buffer
	if err := t.templates.ExecuteTemplate(&tpl, t.open_tmpl, info); err != nil {
		panic(err)
	}
	return tpl.String()
}

func formatVersicles(t *FormatTemplates, cfg Configuration, bookName string, chapterNumber string, versicleNumber string, versicleText string) string {
	info := Info{cfg.GetOuterSeparator(), cfg.GetMajorSeparator(), cfg.GetMinorSeparator(), bookName, chapterNumber, versicleNumber, template.HTML(versicleText)}
	var tpl bytes.Buffer
	if err := t.templates.ExecuteTemplate(&tpl, t.versicle_tmpl, info); err != nil {
		panic(err)
	}
	return tpl.String()
}

func formatClose(t *FormatTemplates, cfg Configuration, bookName string, chapterNumber string, versicleNumber string, versicleText string) string {
	info := Info{cfg.GetOuterSeparator(), cfg.GetMajorSeparator(), cfg.GetMinorSeparator(), bookName, chapterNumber, versicleNumber, template.HTML(versicleText)}
	var tpl bytes.Buffer
	if err := t.templates.ExecuteTemplate(&tpl, t.close_tmpl, info); err != nil {
		panic(err)
	}
	return tpl.String()
}

func TestamentsSearchOpenFormat(cfg Configuration) string {
	return formatOpen(testamentsSearchTemplates, cfg, "", "", "", "")
}

func BookSearchOpenFormat(cfg Configuration, bookName string) string {
	return formatOpen(bookSearchTemplates, cfg, bookName, "", "", "")
}

func ChapterSearchOpenFormat(cfg Configuration, bookName string, chapterNumber string) string {
	return formatOpen(chapterSearchTemplates, cfg, bookName, chapterNumber, "", "")
}

func ReadOpenFormat(cfg Configuration, bookName string) string {
	return formatOpen(readTemplates, cfg, bookName, "", "", "")
}

func TestamentsSearchVersiclesFormat(cfg Configuration, bookName string, chapterNumber string, versicleNumber string, versicleText string) string {
	return formatVersicles(testamentsSearchTemplates, cfg, bookName, chapterNumber, versicleNumber, versicleText)
}

func BookSearchVersiclesFormat(cfg Configuration, bookName string, chapterNumber string, versicleNumber string, versicleText string) string {
	return formatVersicles(bookSearchTemplates, cfg, bookName, chapterNumber, versicleNumber, versicleText)
}

func ChapterSearchVersiclesFormat(cfg Configuration, bookName string, chapterNumber string, versicleNumber string, versicleText string) string {
	return formatVersicles(chapterSearchTemplates, cfg, bookName, chapterNumber, versicleNumber, versicleText)
}

func ReadVersiclesFormat(cfg Configuration, bookName string, chapterNumber string, versicleNumber string, versicleText string) string {
	return formatVersicles(readTemplates, cfg, bookName, chapterNumber, versicleNumber, versicleText)
}

func TestamentsSearchCloseFormat(cfg Configuration) string {
	return formatClose(testamentsSearchTemplates, cfg, "", "", "", "")
}

func BookSearchCloseFormat(cfg Configuration, bookName string) string {
	return formatClose(bookSearchTemplates, cfg, bookName, "", "", "")
}

func ChapterSearchCloseFormat(cfg Configuration, bookName string, chapterNumber string) string {
	return formatClose(chapterSearchTemplates, cfg, bookName, chapterNumber, "", "")
}

func ReadCloseFormat(cfg Configuration, bookName string) string {
	return formatClose(readTemplates, cfg, bookName, "", "", "")
}
