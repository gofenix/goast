package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

func getAllGoFiles() []string {
	dir := "./demo"

	paths := []string{}
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println("遍历出错", err)
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("遍历出错", err)
	}

	return paths
}

func getFunctionCount() {
	paths := getAllGoFiles()
	fset := token.NewFileSet()

	functionCount := 0
	for _, p := range paths {
		f, err := parser.ParseFile(fset, p, nil, 0)
		if err != nil {
			continue
		}

		ast.Inspect(f, func(n ast.Node) bool {
			if _, ok := n.(*ast.FuncDecl); ok {
				functionCount++
			}
			return true
		})
	}
	fmt.Println("function count: ", functionCount)
}

func getFunctionLines() {
	paths := getAllGoFiles()
	fset := token.NewFileSet()

	for _, p := range paths {
		f, err := parser.ParseFile(fset, p, nil, 0)
		if err != nil {
			continue
		}

		ast.Inspect(f, func(n ast.Node) bool {
			if decl, ok := n.(*ast.FuncDecl); ok {
				functionName := decl.Name.Name
				lineCount := fset.Position(decl.End()).Line - fset.Position(decl.Pos()).Line
				// 打印函数名和代码行数
				fmt.Printf("函数名: %s, 代码行数: %d\n", functionName, lineCount)
			}
			return true
		})
	}
}

func getGoMod() {
	modFilePath := "go.mod"

	data, err := ioutil.ReadFile(modFilePath)
	if err != nil {
		log.Fatal("解析Go mod文件错误:", err)
	}

	modFile, err := modfile.Parse(modFilePath, data, nil)
	if err != nil {
		log.Fatal("解析Go mod文件错误:", err)
	}
	fmt.Println("Go version: ", modFile.Go.Version)
	fmt.Println("Go module: ", modFile.Module.Mod.Path)
	fmt.Println("外部依赖数：", len(modFile.Require))
}

func main() {
	getGoMod()
	getFunctionCount()
	getFunctionLines()
}
