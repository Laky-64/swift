package main

import (
	"bytes"
	"fmt"
	"go/format"
	"golang.org/x/tools/go/packages"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	baseUrl  = "https://raw.githubusercontent.com/swiftlang/github.com/Laky-64/goswift/main/include/github.com/Laky-64/goswift/Demangling/"
	kindName = "NodeKind"
)

var (
	nodeRgx = regexp.MustCompile(`^(NODE|CONTEXT_NODE)\((.*?)\)`)
	stdRgx  = regexp.MustCompile(`^(STANDARD_TYPE|STANDARD_TYPE_CONCURRENCY)\((.*?), (.*?), (.*?)\)`)
)

func getContent(fileName string) []string {
	res, err := http.Get(fmt.Sprintf("%s%s", baseUrl, fileName))
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	for {
		buf := make([]byte, 4096)
		n, err := res.Body.Read(buf)
		if n > 0 {
			data = append(data, buf[:n]...)
		}
		if err != nil {
			break
		}
	}
	err = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}
func usize(n int) int {
	switch {
	case n < 1<<8:
		return 8
	case n < 1<<16:
		return 16
	default:
		// 2^32 is enough constants for anyone.
		return 32
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("generator: ")
	pkgList, err := packages.Load(&packages.Config{})
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	buf.WriteString("// Code autogenerated; DO NOT EDIT.\n\n")
	buf.WriteString(fmt.Sprintf("package %s\n\n", pkgList[0].Name))
	buf.WriteString("import \"fmt\"\n")
	buf.WriteString(fmt.Sprintf("type %s int\n\n", kindName))
	var constants []string
	for _, line := range getContent("DemangleNodes.def") {
		if matches := nodeRgx.FindAllStringSubmatch(line, -1); len(matches) > 0 {
			constants = append(constants, matches[0][2])
		}
	}
	buf.WriteString("const (\n")
	for i, constant := range constants {
		buf.WriteString(fmt.Sprintf("\t%sKind", constant))
		if i == 0 {
			buf.WriteString(fmt.Sprintf(" %s = iota", kindName))
		}
		buf.WriteString("\n")
	}
	buf.WriteString(")\n\n")

	buf.WriteString("func _() {\n")
	buf.WriteString("\t// An \"invalid array index\" compiler error signifies that the constant values have changed.\n")
	buf.WriteString("\t// Re-run the generator to resolve the issue.\n")
	buf.WriteString("\tvar x [1]struct{}\n")
	for i, constant := range constants {
		buf.WriteString(fmt.Sprintf("\t_ = x[%sKind-%d]\n", constant, i))
	}
	buf.WriteString("}\n\n")

	var b bytes.Buffer
	indexes := make([]int, len(constants))
	for i := range constants {
		b.WriteString(constants[i])
		indexes[i] = b.Len()
	}
	buf.WriteString(fmt.Sprintf("const _%sName = %q\n", kindName, b.String()))
	nameLen := b.Len()
	buf.WriteString(fmt.Sprintf("var _%sIndex = [...]uint%d{0, ", kindName, usize(nameLen)))
	for i, v := range indexes {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("%d", v))
	}
	buf.WriteString("}\n\n")
	buf.WriteString(fmt.Sprintf("func (i %s) String() string {\n", kindName))
	buf.WriteString(fmt.Sprintf("\tif i >= %s(len(_%sIndex)-1) {\n", kindName, kindName))
	buf.WriteString(fmt.Sprintf("\t\treturn fmt.Sprintf(\"%s(%%d)\", i)\n", kindName))
	buf.WriteString("\t}\n")
	buf.WriteString(fmt.Sprintf("\treturn _%sName[_%sIndex[i]:_%sIndex[i+1]]\n", kindName, kindName, kindName))
	buf.WriteString("}\n\n")

	buf.WriteString("type StandardType struct {\n")
	buf.WriteString(fmt.Sprintf("\tKind     %s\n", kindName))
	buf.WriteString("\tTypeName string\n")
	buf.WriteString("}\n\n")
	var bufStd bytes.Buffer
	var bufStdC bytes.Buffer
	for _, line := range getContent("StandardTypesMangling.def") {
		if matches := stdRgx.FindAllStringSubmatch(line, -1); len(matches) > 0 {
			l := fmt.Sprintf("\t'%s': {%sKind, \"%s\"},\n", matches[0][3], matches[0][2], matches[0][4])
			if matches[0][1] == "STANDARD_TYPE" {
				bufStd.WriteString(l)
			} else {
				bufStdC.WriteString(l)
			}
		}
	}
	buf.WriteString("var standardTypes = map[rune]StandardType{\n")
	buf.WriteString(bufStd.String())
	buf.WriteString("}\n\n")
	buf.WriteString("var standardTypesConcurrency = map[rune]StandardType{\n")
	buf.WriteString(bufStdC.String())
	buf.WriteString("}\n")
	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("standard_types.go", src, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
