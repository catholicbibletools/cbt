package bibles

type Abbreviations = map[string][]string

var abbreviations Abbreviations = map[string][]string{
	"VPGN": {"gn", "ge", "gen", "gén", "gên", "1mo", "imo"},
	"VPEX": {"ex", "exo", "exod", "éx", "éxo", "éxod", "êx", "êxo", "êxod", "eso", "esod", "2mo", "iimo"},
	"VPLV": {"lv", "le", "lev", "3mo", "iiimo"},
	"VPNM": {"nm", "nu", "nb", "num", "núm", "4mo", "ivmo"},
	"VPDT": {"dt", "de", "deu", "deut", "5mo", "vmo"},
	"VHIE": {"js", "jos", "gs", "ios", "jsh", "josh"},
	"VHID": {"jue", "idc", "gdc", "judg", "jdg", "jg", "jdgs", "jz", "jzs", "jg", "jug", "ri"},
	"VH1S": {"1s", "1sa", "1sam", "1sm", "ism"},
	"VH2S": {"2s", "2sa", "2sam", "2sm", "iism"},
	"VH1R": {"1r", "1rs", "1re", "1rois", "1king", "1kgs", "1kin", "1ki", "1kg", "1kön", "1reg", "1rg", "irg"},
	"VH2R": {"2r", "2rs", "2re", "2rois", "2king", "2kgs", "2kin", "2ki", "2kg", "2kön", "2reg", "2rg", "iirg"},
	"VH1P": {"1cr", "1cro", "1cron", "1crón", "1chro", "1chr", "1ch", "1par", "ipar", "ichr"},
	"VH2P": {"2cr", "2cro", "2cron", "2crón", "2chro", "2chr", "2ch", "2par", "iipar", "iichr"},
	"VHED": {"ed", "esd", "ezra", "ezr", "esdr", "esr", "esra", "1esd", "1esr", "iesd", "iesr"},
	"VHNH": {"ne", "né", "nee", "neh", "2esd", "2esr", "iiesd", "iiesr"},
	"VH1M": {"1m", "1mc", "1mac", "1macc", "1makk", "1mak", "1ma", "1mc", "1mcc", "imcc"},
	"VH2M": {"2m", "2mc", "2mac", "2macc", "2makk", "2mak", "2ma", "2mc", "2mcc", "iimcc"},
	"VMIS": {"is", "isa", "és", "jes"},
	"VMIE": {"jr", "je", "jé", "jer", "ger", "ier"},
	"VMEZ": {"éz", "ézé", "eze", "ezk", "ezek", "ezeq", "hes", "hese"},
	"VNOS": {"os", "ho", "hos", "ose", "osee", "osée", "osea", "oseas"},
	"VNIL": {"jl", "joe", "joel", "joël", "ioel", "ioe", "il"},
	"VNAM": {"am", "amos", "amós"},
	"VNAB": {"ab", "abd", "ob", "obd", "obad", "abdia"},
	"VNIN": {"gi", "giona", "ion", "jon", "jona", "jonah", "jnh", "jonas"},
	"VNMI": {"mi", "mq", "miq", "mic", "mica", "micah", "mich"},
	"VNNH": {"na", "nah", "naum", "nahum"},
	"VNHA": {"ha", "hab", "aba", "abac"},
	"VNSO": {"sf", "sof", "soph", "zeph", "zep", "zp", "zef", "zf", "ze"},
	"VNAG": {"ag", "agg", "aggeo", "ageu", "ageo", "hag", "hg"},
	"VNZA": {"za", "zc", "zac", "zach", "zech", "zec", "zacc", "sach"},
	"VNML": {"ml", "mal"},
	"VMDN": {"dn", "da", "dan"},
	"VSIB": {"jb", "jó", "job", "gb", "ib", "iob", "hi", "hiob"},
	"VSPR": {"pr", "pv", "pro", "prv", "prov", "spr"},
	"VSEE": {"qo", "qoh", "coe", "coel", "ecle", "eccle", "ecc1es", "koh", "pre", "pred"},
	"VSCC": {"ct", "ca", "can", "cân", "cnt", "cant", "coc", "song", "sos", "hld", "hoh"},
	"VHRT": {"rt", "ru", "rut", "rth", "ruth", "rute"},
	"VMLM": {"lm", "la", "lam", "thr", "klgl"},
	"VHET": {"et", "est", "esth", "ester"},
	"VHTB": {"tb", "tob", "thb", "thob", "tobi", "tobie", "tobit"},
	"VHIH": {"jt", "jdt", "jth", "jdth", "gdt", "idth", "idt"},
	"VMBR": {"ba", "br", "bar", "brch", "brc", "baruc"},
	"VSSP": {"sb", "sab", "sap", "sag", "wis", "ws", "we", "wsh", "weish"},
	"VSEI": {"si", "sir", "eclo", "ecclus", "ecclu", "eccli", "sirac"},
	"VSPS": {"sl", "sal", "salmi", "ps", "psalm", "pslm", "psa", "psm", "pss"},
	"NEMT": {"mt", "ma", "mat", "matt"},
	"NEMC": {"mk", "mr", "mac", "mar", "mrk", "marc", "mark", "marco"},
	"NELC": {"lc", "lu", "luc", "luca", "lk", "luk", "luke"},
	"NEIO": {"io", "gv", "jhn", "joh", "john", "jean", "joão", "juan"},
	"NTAA": {"ac", "act", "acts", "at", "att", "atti", "atos", "hc", "hch", "apg"},
	"NPRM": {"rm", "ro", "rom", "röm"},
	"NP1C": {"1co", "1cor", "1kor", "icor"},
	"NP2C": {"2co", "2cor", "2kor", "iicor"},
	"NPGA": {"ga", "gal", "gál", "gàl", "glt"},
	"NPEP": {"ef", "efe", "ep", "eph", "ephe", "efes"},
	"NPPP": {"fp", "flp", "filp", "phl", "phlp", "php", "pp"},
	"NPCL": {"cl", "co", "col", "cls", "kol", "kls"},
	"NPPM": {"fm", "flm", "film", "phm", "phlm", "phm", "pm"},
	"NP1E": {"1ts", "1te", "1tes", "1tess", "1th", "1the", "1thes", "1thess", "ithe", "ith"},
	"NP2E": {"2ts", "2te", "2tes", "2tess", "2th", "2the", "2thes", "2thess", "iithe", "iith"},
	"NP1I": {"1ti", "1tm", "1tim", "itim", "itm", "iti"},
	"NP2I": {"2ti", "2tm", "2tim", "iitim", "iitm", "iiti"},
	"NPTT": {"tt", "ti", "tit", "tite", "tito", "titus"},
	"NTHB": {"eb", "ebr", "ebrei", "hebr", "hé", "heb", "hbr", "hbrw"},
	"NCIB": {"ia", "iac", "jac", "jak", "gc", "tg", "tgo", "sgo", "sant", "jm", "jms", "jas", "james"},
	"NC1P": {"1p", "1pi", "1pd", "1pe", "1pt", "1pet", "1petr", "1ptr", "ipet", "ipt"},
	"NC2P": {"2p", "2pi", "2pd", "2pe", "2pt", "2pet", "2petr", "2ptr", "iipet", "iipt"},
	"NCID": {"jd", "jds", "jude", "jud", "gd", "giuda", "iud"},
	"NC1I": {"1j", "1jn", "1jhn", "1jo", "1joh", "1gv", "1io", "iio"},
	"NC2I": {"2j", "2jn", "2jhn", "2jo", "2joh", "2gv", "2io", "iiio"},
	"NC3I": {"3j", "3jn", "3jhn", "3jo", "3joh", "3gv", "3io", "iiiio"},
	"NTAI": {"ap", "apc", "apk", "apoc", "apok", "re", "rv", "rev", "of", "off", "offb"},
}

func newAbbreviations(merge *Abbreviations) *Abbreviations {
	result := Abbreviations{}
	for code := range abbreviations {
		result[code] = abbreviations[code]
	}
	for code := range *merge {
		abbrvs := (*merge)[code]
		for index := range abbrvs {
			abbr := abbrvs[index]
			if !isPresent(result[code], abbr) {
				result[code] = append(result[code], abbr)
			}
		}
	}
	return &result
}

var abbreviations_ES Abbreviations = map[string][]string{
	"VHID": {"jc"},
	"VHED": {"es"},
	"VMEZ": {"ez"},
	"VNIN": {"jo"},
	"VNSO": {"so"},
	"NEMC": {"mc"},
	"NEIO": {"jn"},
	"NTAA": {"he"},
	"NTHB": {"hb"},
	"NCIB": {"sg"},
}

var abbreviations_EN Abbreviations = map[string][]string{
	"VHED": {"ez"},
	"VNMI": {"mc"},
	"VNHA": {"hb"},
	"VSCC": {"so"},
	"VHET": {"es"},
	"NEIO": {"jn"},
	"NPGA": {"gl"},
}

var abbreviations_PT Abbreviations = map[string][]string{
	"VMEZ": {"ez"},
	"VNIN": {"jn"},
	"NEMC": {"mc"},
	"NEIO": {"jo"},
	"NPGA": {"gl"},
	"NTHB": {"hb"},
}

var abbreviations_DE Abbreviations = map[string][]string{
	"VMEZ": {"ez"},
}

var abbreviations_FR Abbreviations = map[string][]string{
	"VMIS": {"es"},
	"VMEZ": {"ez"},
	"VNSO": {"so"},
	"VSSP": {"sg"},
	"NEMC": {"mc"},
	"NEIO": {"jn"},
	"NTHB": {"he"},
	"NCIB": {"jc"},
}

var abbreviations_IT Abbreviations = map[string][]string{
	"VPEX": {"es"},
	"VMEZ": {"ez"},
	"VNIL": {"gl"},
	"VNSO": {"so"},
	"NEMC": {"mc"},
}

var abbreviations_LA Abbreviations = map[string][]string{
	"VMEZ": {"ez"},
	"VNSO": {"so"},
	"NEMC": {"mc"},
}

/*func hasIndex(array []string, index int) bool {
	return (len(array) > index)
}*/

func isPresent(array []string, value string) bool {
	for _, b := range array {
		if b == value {
			return true
		}
	}
	return false
}

func GetBookCodeByAbbreviation(abbreviations *Abbreviations, abbreviation string) string {
	for code := range *abbreviations {
		if isPresent((*abbreviations)[code], abbreviation) {
			return code
		}
	}
	return ""
}

func GetAbbreviations(language string) *Abbreviations {
	switch language {
	case "es":
		return newAbbreviations(&abbreviations_ES)
	case "pt":
		return newAbbreviations(&abbreviations_PT)
	case "de":
		return newAbbreviations(&abbreviations_DE)
	case "fr":
		return newAbbreviations(&abbreviations_FR)
	case "it":
		return newAbbreviations(&abbreviations_IT)
	case "la":
		return newAbbreviations(&abbreviations_LA)
	}
	return newAbbreviations(&abbreviations_EN)
}

func IsValidLanguage(language string) bool {
	switch language {
	case "es", "pt", "de", "fr", "it", "la", "en":
		return true
	}
	return false
}
