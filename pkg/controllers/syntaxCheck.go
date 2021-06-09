package controllers

import (
	"app/diplom/pkg/models"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func SyntaxCheck(taskDecision string) (*models.AnswerTask, error) {
	var answer models.AnswerTask
	var answerString string
	answerString, err := syntaxCheck(taskDecision)
	if err != nil {
		answer.Answer = "error"
		return &answer, err
	}
	answer.Answer = answerString
	return &answer, nil
}

func syntaxCheck(taskDecision string)(answer string, err error){
	myErr := ""
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", taskDecision, 0)
	if err != nil {
		myErr = err.Error()
		newErr := myErr[:strings.IndexByte(myErr, ':')]
		answer += "Ошибка в "
		answer += newErr
		answer += " строке"
		answer += "\n"
	} else {
		ast.Print(fset, f)

		var packageName string

		ast.Inspect(f, func(x ast.Node) bool {
			s, ok := x.(*ast.File)
			if !ok {
				return true
			}
			if s.Name.Name == "main"{
				packageName = "M"
				answer += "S  ->  pM"
				answer += "\n"
				fmt.Println("S -> pM")
			} else{
				packageName = s.Name.Name
				fmt.Println("S  ->  p" + s.Name.Name)
				answer += "S  ->  p"
				answer += s.Name.Name
				answer += "\n"
			}

			return false
		})
		ast.Inspect(f, func(x ast.Node) bool {
			s, ok := x.(*ast.FuncDecl)
			if !ok {
				return true
			}
			if s.Name.Obj.Kind.String() == "func" && s.Name.Obj.Name == "main"{
				fmt.Println(packageName + "  -> fm(){N}")
				answer += packageName
				answer += "  ->  fm(){N}"
				answer += "\n"

			} else{
				fmt.Println(packageName + " -> fl(){N}")
				answer += packageName
				answer += "  ->  fl(){N}"
				answer += "\n"
			}

			return false
		})
		ast.Inspect(f, func(x ast.Node) bool {
			//N- i(F)
			str := "N  ->  "
			s, ok := x.(*ast.CallExpr)
			if !ok {
				return true
			}

			ast.Inspect(s, func(x ast.Node) bool {
				_, ok := x.(*ast.Ident)
				if !ok {
					return true
				}
				str += "i(F)"
				return false
			})
			fmt.Println(str)
			answer += str
			answer += "\n"
			ast.Inspect(s, func(x ast.Node) bool {
				l, ok := x.(*ast.BasicLit)
				if !ok {
					return true
				}
				if l.Kind == token.STRING || l.Kind == token.INT{
					fmt.Println("F -> l")
					answer += "F  ->  l"
					answer += "\n"
				}
				return false
			})

			return false
		})
	}
	return answer, nil
}
