## Generate codelab

Install claat with `go get -u github.com/googlecodelabs/tools/claat`

At repository root, execute `claat export -o docs/ codelab.md`

## Generate index

Install mdtohtml with `go get -u github.com/gomarkdown/mdtohtml`

At repository root, execute `mdtohtml -page -xhtml=false -css https://nlepage.github.io/catption/index.css docs/index.md docs/index.html`
