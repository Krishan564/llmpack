package skeleton

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
)

func reduceGo(content []byte) ([]byte, error) {
	fset := token.NewFileSet()
	// ParseComments важливо, щоб зберегти документацію до функцій
	node, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		return nil, err // Або повернути content, якщо хочеш soft fail
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// Шукаємо оголошення функцій (func main, func (s *Struct) Method)
		if fn, ok := n.(*ast.FuncDecl); ok {
			// Якщо у функції є тіло (вона не інтерфейс і не forward declaration)
			if fn.Body != nil {
				// Замінюємо список інструкцій на порожній список або список з одним коментарем
				fn.Body.List = []ast.Stmt{
					&ast.ExprStmt{
						X: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "`... implementation hidden ...`", // Або просто коментар
						},
					},
				}
			}
		}
		return true
	})

	// Рендеримо AST назад у код
	var buf bytes.Buffer
	// printer.TabIndent зберігає оригінальне форматування
	if err := printer.Fprint(&buf, fset, node); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
