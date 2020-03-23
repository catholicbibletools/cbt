package main

import (
	"os"
	"path/filepath"
	"strings"

	. "../bibles"
)

const (
	TITLE       = "CATHOLiC BiBLE TOOLS"
	VERSION     = "1.00"
	CONFIG_FILE = "cbt.cfg"
	//BIBLES_DIR  = "bibles"
)

var (
	currentBible         *Bible
	currentAbbreviations *Abbreviations
	currentConfiguration Configuration
)

func initialize(path string) {
	dir, err1 := filepath.Abs(filepath.Dir(path))
	if err1 != nil {
		panic(err1)
	}
	err2 := os.Chdir(dir)
	if err2 != nil {
		panic(err2)
	}
	LoadBiblesFromResources()
	//LoadBiblesFromDirectory(BIBLES_DIR)
	currentConfiguration = LoadConfigurationFromFile(CONFIG_FILE)
	LoadTemplates(currentConfiguration.GetOutputFormat())
}

func main() {
	initialize(os.Args[0])
	la := len(os.Args)
	var result *CommandResult
	if la > 1 {
		switch os.Args[1] {
		case "v", "version":
			result = CommandVersion()
		case "l", "list":
			result = CommandList()
		case "r", "read":
			if la > 2 {
				currentBible = GetBible(currentConfiguration.GetCurrentBible())
				currentBible.LoadBooks()
				currentAbbreviations = GetAbbreviations(currentConfiguration.GetCurrentLanguage())
				result = CommandRead(strings.Join(os.Args[2:], " "))
			}
		case "s", "search":
			if la > 2 {
				currentBible = GetBible(currentConfiguration.GetCurrentBible())
				currentBible.LoadBooks()
				currentAbbreviations = GetAbbreviations(currentConfiguration.GetCurrentLanguage())
				where := ""
				excluding := ""
				if la > 3 {
					if la > 4 {
						if strings.HasPrefix(os.Args[3], "-") {
							excluding = strings.TrimPrefix(os.Args[3], "-")
							where = strings.Join(os.Args[4:], " ")
						} else {
							where = strings.Join(os.Args[3:], " ")
						}
					} else {
						where = strings.Join(os.Args[3:], " ")
					}
				}
				result = CommandSearch(os.Args[2], excluding, where)
			}
		case "w", "web":
			if la > 2 {
				currentBible = GetBible(currentConfiguration.GetCurrentBible())
				currentBible.LoadBooks()
				currentAbbreviations = GetAbbreviations(currentConfiguration.GetCurrentLanguage())
				result = CommandWeb(os.Args[2])
			}
		case "c", "config":
			if la > 2 {
				if la > 3 {
					result = CommandConfig(os.Args[2], os.Args[3])
				} else {
					result = CommandConfig(os.Args[2], "")
				}
			} else {
				result = CommandConfig("", "")
			}
		case "h", "help":
			if la > 2 {
				result = CommandHelp(os.Args[2])
			}
		}
	}
	if result == nil {
		result = CommandHelp("")
	}
	os.Exit(ProcessCommandResult(result))
}
