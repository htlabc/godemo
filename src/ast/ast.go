package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func AstDemo() {

	// 这是我们需要进行语法树分析的代码
	src := `
package p
const c = 1.0
var X = f(3.14)*2 + c
`

	// 创建一个AST
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	// 遍历语法树节点
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
		case *ast.Ident:
			s = x.Name
		}
		if s != "" {
			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})

}
