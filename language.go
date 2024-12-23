package clocexplorer

import (
	"path/filepath"
)

type Language struct {
	Name         string
	lineComments []string
	Files        []string
	Code         uint64
	Comments     uint64
	Blanks       uint64
}

var Exts = map[string]string{
	"go":  "Go",
	"c":   "C",
	"cpp": "C++",
	"rs":  "Rust",
}

func GetFileType(path string) (ext string, ok bool) {
	ext = filepath.Ext(path)

	for k, v := range Exts {
		if ext == "" {
			continue
		}

		if k == ext[1:] {
			return v, true
		}
	}
	return "", false
}

func NewLanguage(name string, lineComments []string) *Language {
	return &Language{
		Name:         name,
		lineComments: lineComments,
		Files:        []string{},
	}
}

type DefinedLanguages struct {
	Langs map[string]*Language
}

func NewDefinedLanguages() *DefinedLanguages {
	return &DefinedLanguages{
		Langs: map[string]*Language{
			"Go":   NewLanguage("Go", []string{"//"}),
			"C":    NewLanguage("C", []string{"//"}),
			"C++":  NewLanguage("C++", []string{"//"}),
			"Rust": NewLanguage("Rust", []string{"//"}),
		},
	}
}
