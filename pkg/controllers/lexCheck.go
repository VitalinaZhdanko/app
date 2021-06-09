package controllers

import (
	"app/diplom/pkg/models"
	"fmt"
	"go/parser"
	"go/scanner"
	"go/token"
	"strconv"
	"strings"
)

func LexCheckByTaskID(taskDecision string) (*models.AnswerTask, error) {
	var answer models.AnswerTask
	var answerString string
	answerString, err := lexCheck(taskDecision)
	if err != nil {
		answer.Answer = "error"
		return &answer, err
	}
	answer.Answer = answerString
	return &answer, nil
}

func lexCheck(src string) (answer string, err error){
	myErr := ""
	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, "", src, 0)
	if err != nil {
		myErr = err.Error()
		fmt.Println(myErr)
		newErr := myErr[:strings.IndexByte(myErr, ':')]
		fmt.Println(newErr)
		answer += "Ошибка в "
		answer += newErr
		answer += " строке"

	} else {

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, []byte(src), nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		if tok == token.SEMICOLON {
			continue
		}
		if tok == token.PACKAGE{
			fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), "p", lit)
			column := strconv.Itoa(fset.Position(pos).Column)
			line := strconv.Itoa(fset.Position(pos).Line)
			answer += line
			answer += ":"
			answer += column
			answer += "\t"
			answer += "p"
			answer += "\t"
			answer += lit
			answer += "\n"
			continue
		}

		if tok == token.IDENT && lit == "main"{
			fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), "m", lit)
			column := strconv.Itoa(fset.Position(pos).Column)
			line := strconv.Itoa(fset.Position(pos).Line)
			answer += line
			answer += ":"
			answer += column
			answer += "\t"
			answer += "m"
			answer += "\t"
			answer += lit
			answer += "\n"
			continue
		}
		if tok == token.FUNC{
			fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), "f", lit)
			column := strconv.Itoa(fset.Position(pos).Column)
			line := strconv.Itoa(fset.Position(pos).Line)
			answer += line
			answer += ":"
			answer += column
			answer += "\t"
			answer += "f"
			answer += "\t"
			answer += lit
			answer += "\n"
			continue
		}

		if tok == token.IDENT || tok == token.VAR || tok == token.RETURN{
			fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), "i", lit)
			column := strconv.Itoa(fset.Position(pos).Column)
			line := strconv.Itoa(fset.Position(pos).Line)
			answer += line
			answer += ":"
			answer += column
			answer += "\t"
			answer += "i"
			answer += "\t"
			answer += lit
			answer += "\n"
			continue
		}
		if tok == token.STRING{
			fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), "l", lit)
			column := strconv.Itoa(fset.Position(pos).Column)
			line := strconv.Itoa(fset.Position(pos).Line)
			answer += line
			answer += ":"
			answer += column
			answer += "\t"
			answer += "l"
			answer += "\t"
			answer += lit
			answer += "\n"
			continue
		}
		if tok == token.LPAREN || tok == token.RPAREN || tok == token.LBRACE || tok == token.RBRACE ||
			tok == token.ADD || tok == token.SUB || tok == token.MUL || tok == token.QUO || tok == token.REM ||
			tok == token.AND || tok == token.OR || tok == token.XOR || tok == token.SHL || tok == token.SHR ||
			tok == token.AND_NOT || tok == token.ADD_ASSIGN || tok == token.SUB_ASSIGN || tok == token.MUL_ASSIGN ||
			tok == token.QUO_ASSIGN || tok == token.REM_ASSIGN || tok == token.AND_ASSIGN || tok == token.OR_ASSIGN ||
			tok == token.XOR_ASSIGN || tok == token.SHL_ASSIGN || tok == token.SHR_ASSIGN || tok == token.AND_NOT_ASSIGN ||
			tok == token.LAND || tok == token.LOR || tok == token.ARROW || tok == token.INC || tok == token.DEC ||
			tok == token.EQL || tok == token.LSS || tok == token.GTR || tok == token.ASSIGN || tok == token.NOT ||
			tok == token.NEQ || tok == token.LEQ || tok == token.GEQ || tok == token.DEFINE || tok == token.ELLIPSIS ||
			tok == token.COMMA{
			fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, tok)
			column := strconv.Itoa(fset.Position(pos).Column)
			line := strconv.Itoa(fset.Position(pos).Line)
			answer += line
			answer += ":"
			answer += column
			answer += "\t"
			answer += tok.String()
			answer += "\t"
			answer += tok.String()
			answer += "\n"
			continue
		}

		if tok == token.INT || tok == token.FLOAT || tok == token.CHAR{
			fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), "l", lit)
			column := strconv.Itoa(fset.Position(pos).Column)
			line := strconv.Itoa(fset.Position(pos).Line)
			answer += line
			answer += ":"
			answer += column
			answer += "\t"
			answer += "l"
			answer += "\t"
			answer += lit
			answer += "\n"
			continue
		}

		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
		column := strconv.Itoa(fset.Position(pos).Column)
		line := strconv.Itoa(fset.Position(pos).Line)
		answer += line
		answer += ":"
		answer += column
		answer += "\t"
		answer += tok.String()
		answer += "\t"
		answer += lit
		answer += "\n"
	}
	}
	return answer, nil
}