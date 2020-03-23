# TOWARDS 2.0

---

## MUST HAVE

### CLIENT
- a minimal fully featured client

### BIBLES
- add it.sb from Vatican site: http://www.vatican.va/archive/ITA0001/_INDEX.HTM
- add en.nab from Vatican site: http://www.vatican.va/archive/ENG0839/_INDEX.HTM

---

## COULD HAVE

### OUTPUT
- add support to output-format for:
    tsv -- https://en.wikipedia.org/wiki/Tab-separated_values
    yaml -- https://en.wikipedia.org/wiki/YAML
    tml -- https://en.wikipedia.org/wiki/TOML
    ini -- https://en.wikipedia.org/wiki/INI_file

---

## SHOULD HAVE

### INTERNATIONALIZATION
- use current-language for interface messages

---

## WONT HAVE

### BOOKMARKS
- add bookmark cli command and web resource
    b -- list bookmarks for current bible
    b $title $reference -- bookmark for current bible
    POST /bookmark
    GET /bookmarks
    file cbt_bookmarks.cfg in bible: *[{ "title": "Apocalipsis 1", "reference": "ap 1" }]*

### READ
- add switch to allow reading a passage in all of the bibles for the current language -- it's like a compare feature

### SEARCH
- *feature:* add mixed versicle hit (end of one versicle plus the beginning of the next one are the place where the input hits)
- *enhancement:* add BibleSearcher with callbacks to determine BookCode searchability -ex. if it starts with N for a new testament search, etc- and handle hit deciding if add to content and how it will be added

### SEPARATORS
- add union citation separator -- check: https://www.eltestigofiel.org/index.php?idu=pb_10990 (and http://usitep.es/apf/reli/citas_biblicas/citas-signos.html)
    "union-separator": "." // indicates chapters or versicles that are not part of a range but are contigous (ex. Gn 1,1.26; Jn 2,3.7; etc)