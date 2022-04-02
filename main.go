package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) == 2 {
		check(os.Chdir(os.Args[1]))
	}
	b, err := exec.Command("go", "list", "-f", "{{.Dir}}", "./...").CombinedOutput()
	check(err)
	dirs := strings.Split(strings.TrimSpace(string(b)), "\n")
	fns := map[string]int{}
	for _, dir := range dirs {
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, dir, nil, 0)
		check(err)
		for name, pkg := range pkgs {
			if strings.HasSuffix(name, "_test") {
				continue
			}
			for name, file := range pkg.Files {
				if strings.HasSuffix(name, "_test.go") {
					continue
				}
				for _, decl := range file.Decls {
					if fn, ok := decl.(*ast.FuncDecl); ok {
						if fn.Name.IsExported() {
							fns[fn.Name.Name]++
						}
					}
				}
			}
		}
	}
	var total int
	var names []string
	for k, v := range fns {
		names = append(names, k)
		total += v
	}
	width := len(strconv.Itoa(total))
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%*d %s\n", width, fns[name], name)
	}
	fmt.Println(total, "exported functions")
}
