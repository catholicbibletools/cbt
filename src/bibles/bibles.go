package bibles

/*import "io/ioutil"

func listFolders(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		if file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

func LoadBiblesFromDirectory(path string) {
	files, err := listFolders(path)
	if err == nil {
		for n := range files {
			Bibles = append(Bibles, *NewBible(path, files[n]))
		}
	}
}*/

var Bibles []Bible = []Bible{}

func LoadBiblesFromResources() {
	Bibles = append(Bibles, *NewBible( /*"",*/ "en.ccb"))
	Bibles = append(Bibles, *NewBible( /*"",*/ "en.drbcr"))
	Bibles = append(Bibles, *NewBible( /*"",*/ "es.eldpdd"))
	Bibles = append(Bibles, *NewBible( /*"",*/ "es.lbl"))
	Bibles = append(Bibles, *NewBible( /*"",*/ "es.lsb"))
	Bibles = append(Bibles, *NewBible( /*"",*/ "pt.bs"))
	Bibles = append(Bibles, *NewBible( /*"",*/ "fr.lbdcc"))
	Bibles = append(Bibles, *NewBible( /*"",*/ "it.lb"))
	Bibles = append(Bibles, *NewBible( /*"",*/ "la.nvbse"))
}

func GetBible(code string) *Bible {
	for n := range Bibles {
		var b = Bibles[n]
		if b.code == code {
			return &b
		}
	}
	return nil
}
