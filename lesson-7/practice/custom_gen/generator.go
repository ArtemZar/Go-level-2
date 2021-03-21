package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
)

// шаблон для генерированного файла
const constListTmpl = `// CODE GENERATED AUTOMATICALLY
// THIS FILE SHOULD NOT BE EDITED BY HAND
package {{.Package}}

type {{.Name}}s []{{.Name}}

func (c {{.Name}}s) List() []{{.Name}} {
	return []{{.Name}}{{"{"}}{{.List}}{{"}"}}
}
`

const (
pkgName = "main"

srcFileName = "src.go" // название файла, по которому мы итерируемся для составления AST
genFileName = "list_gen.go"

typeName = "Color" // тип констант для которых будет создан список
)

func main() {
consts, err := getConstants(srcFileName, typeName)
if err != nil {
log.Fatal(err)
}

templateData := struct {
Package string
Name    string
List    string
}{
Package: pkgName,
Name:    typeName,
List:    strings.Join(consts, ", "),
}

genFile, err := os.Create(genFileName)
if err != nil {
log.Fatal(err)
}

t := template.Must(template.New("const-list").Parse(constListTmpl))
if err := t.Execute(genFile, templateData); err != nil {
log.Fatal(err)
}
}

func getConstants(srcFileName, typeName string) ([]string, error) {
fset := token.NewFileSet()
// парсим файл, чтобы получить AST
astFile, err := parser.ParseFile(fset, srcFileName, nil, 0)
if err != nil {
return nil, err
}

var (
constType string
out       []string
)

for _, decl := range astFile.Decls {
genDecl, ok := decl.(*ast.GenDecl)
if !ok {
continue
}

// пропускаем все декларации, которые не константы
if genDecl.Tok != token.CONST {
continue
}

// итерируемся по элементам: с определениями типов, переменных, констант, функций и т.п.
for _, spec := range genDecl.Specs {
vspec, ok := spec.(*ast.ValueSpec) // отсюда мы получим наименование константы
if !ok {
continue
}

if vspec.Type == nil && len(vspec.Values) > 0 {
// случай определения константы как "X = 1"
// такая константа не имеет типа и может быть пропущена
// это может означать, что был начат новый блок определения const
constType = ""
continue
}

if vspec.Type != nil {
ident, ok := vspec.Type.(*ast.Ident)
if !ok {
continue
}
// здесь записываем тип константы
constType = ident.Name
}

// при совпадении типа константы с заданным, записываем в выходные параметры
if constType == typeName {
if len(vspec.Names) == 0 {
continue
}

out = append(out, vspec.Names[0].Name)
}
}
}

return out, nil
}
