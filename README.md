# CATHOLiC BiBLE TOOLS

**cbt** is a command line tool for reading and searching a set of catholic bibles which also provides a web api service.
It supports many abbreviation languages for querying the bible text and different output formats to display results.
It is coded in golang, open sourced and MIT licensed.

## Application Commands

**cbt** accepts the following commands:

* `v` -- *version*
* `h` -- *help*
* `l` -- *list*
* `c` -- *config*
* `r` -- *read*, to read a passage or chapter the currently set bible
* `s` -- *search*, to search the currently set bible
* `w` -- *web*, to serve the web json api

### Version Command
#### Desription
  shows the current version
#### Usage
  `cbt v`

### Help Command
#### Desription
  shows help about another command
#### Usage
  `cbt h <command>`

### List Command
#### Desription
  lists the available bibles
#### Usage
  `cbt l`

### Config Command
#### Desription
  displays current configuration settings or sets specific setting to the current configuration
#### Usage
  `cbt c [<setting> [<value>]]`
#### Examples
  `cbt c`
    *displays all the current configuration settings*

  `cbt c current-bible`
    *displays the current bible in the current configuration settings*

  `cbt c current-bible en.ccb`
    *sets the current bible to 'en.ccb' in the current configuration settings*

### Read Command
#### Desription
  reads the bible
#### Usage
  `cbt r <what>`
#### Examples
  `cbt r 7 gn`
    *reads Genesis chapter 7*

  `cbt r 1 jn 1,1`
    *reads 1 John chapter 1 versicle 1*

  `cbt r 2 co 3,1-5`
    *reads 2 Corinthians chapter 2 versicles 1 to 5*

  `cbt r ap 1-3`
    *reads Apokalypsis chapters 1 to 3*

  `cbt r mt 10,1-12,2`
    *reads Matthew from versicle 10,1 to 12,2*

### Search Command
#### Desription
  searches the bible
#### Usage
  `cbt s '<what>' [-'<excluding>'] [<where>]`
#### Examples
  `cbt s 'love is'`
    *searches all the Bible for 'love is'*

  `cbt s 'love is' o`
    *searches the Old Testament for 'love is'*

  `cbt s 'love is' n`
    *searches the New Testament for 'love is'*

  `cbt s 'love is' ap`
    *searches the Apokalypsis for 'love is'*

  `cbt s 'love is' ap 1`
    *searches the Apokalypsis chapter 1 for 'love is'*

  `cbt s 'love is' 1 jn 1`
    *searches 1 John chapter 1 for 'love is'*

  `cbt s 'the first' -'fire' ap`
    *searches the Apokalypsis for 'the first' excluding verses where 'over them' appears*

### Web Command
#### Desription
  starts the web api
#### Usage
  `cbt w <port>`


## Application Configuration

**cbt** stores it's configuration in a json file called *cbt.cfg* located in the same directory of the **cbt** executable.

### Current Bible

The configuration field `current-bible` determines the currently used bible.

It accepts the following values:

* `en.ccb` -- *Christian Community Bible* **(DEFAULT)**
* `en.drbcr` -- *Douay-Rheims Bible (Challoner Revision)*
* `es.eldpdd` -- *El Libro del Pueblo de Dios*
* `es.lbl` -- *La Biblia Latinoamericana*
* `es.lsb` -- *La Santa Biblia*
* `pt.bs` -- *Biblia Sagrada*
* `fr.lbdcc` -- *La Bible Des Communautés Chrétiennes*
* `it.lb` -- *La Bibbia*
* `la.nvbse` -- *Nova Vulgata (Bibliorum Sacrorum Editio)*

### Current Language

The configuration field `current-language` determines the currently used language (for abbreviations only currently).

It accepts the following values:

* `en` -- *english* **(DEFAULT)**
* `es` -- *spanish*
* `pt` -- *portuguese*
* `de` -- *german*
* `fr` -- *french*
* `it` -- *italian*
* `la` -- *latin*

### Output Format

The configuration field `output-format` determines the currently used output format.

It accepts the following values:

* `txt` -- *text* **(DEFAULT)**
* `md` -- *markdown*
* `html` -- *hypertext markup language*
* `xml` -- *extensible markup language*
* `json` -- *javascript object notation*
* `csv` -- *comma separated values*

### Citation Separators

The configuration field `outer-separator` determines the character used to separate different bible citations (ex. Gn 1,2; Ex 3,4; Lv 5,6; etc) and its default value is `;`.

The configuration field `major-separator` determines the character used to separate chapters from versicles (ex. 1,2; 3,4-5; etc) and its default value is "," (note that ":" is also commonly used for this purpose).

The configuration field `minor-separator` determines the character used to separate chapters or versicles (ex. 1-2; 3,4-5; etc) and its default value is "-".

### Example

```json

    {
        "current-bible": "en.ccb",
        "current-language": "en",
        "output-format": "txt",
        "outer-separator": ";",
        "major-separator": ",",
        "minor-separator": "-",
    }

```

## Web JSON API

### Resources

**cbt** web json api supports the following resources:

* **GET** /version -- *gets the **cbt** version*
* **GET** /list -- *gets the list of the available bibles*
* **GET** /config -- *gets the configuration file contents*
* **GET** /config -- *gets an specified configuration key value (sample call body: `{ "setting": "output-format" }`)*
* **POST** /config -- *sets an specified configuration key value (sample call body: `{ "setting": "output-format", "value": "json" }`)*
* **POST** /read -- *performs a bible read (sample call body: `{ "what": "jn 1,1-2" }`)*
* **POST** /search -- *performs a bible search (sample call body: `{ "what": "the first", "excluding": "fire", "where": "ap" }`)*


---

## Version History

* **1.00** -- *Mar 22, 2020*
* **0.15** -- *Mar 21, 2020*
* **0.14** -- *Dec 28, 2019*
* **0.13** -- *Dec 24, 2019*
* **0.12** -- *Dec 08, 2019*
* **0.11** -- *Dec 07, 2019*
* **0.10** -- *Nov 24, 2019*
* **0.09** -- *Nov 23, 2019*
* **0.08** -- *Nov 18, 2019*
* **0.07** -- *Nov 17, 2019*
* **0.06** -- *Nov 16, 2019*
* **0.05** -- *Nov 13, 2019*
* **0.04** -- *Nov 11, 2019*
* **0.03** -- *Nov 10, 2019*
* **0.02** -- *Nov 09, 2019*
* **0.01** -- *Nov 08, 2019*


---

## License (MIT)

>    Copyright (c) 2020 CATHOLiC BiBLE TOOLS.
>
>    Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
>
>    The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
>
>    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
