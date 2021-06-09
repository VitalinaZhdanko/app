package controllers

import (
	"app/diplom/pkg/models"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunTask(taskDecision string) (*models.AnswerTask, error) {
	var answer models.AnswerTask

	answerString, _ := buildUserSolution(taskDecision, "golang")
	fmt.Println("333")
	fmt.Println(answerString)
	//if err != nil {
	//	answer.Answer = "error"
	//	return &answer, err
	//}
	answer.Answer = answerString
	fmt.Println("2222")
	fmt.Println(answer)
	return &answer, nil
}




	//	src := `
	//package main
	//func main() {
	//	println("Hello, World!")
	//}
	//`


//	src2 := `
//#include <iostream>
//using namespace std;
//
//int main() {
//  cout << "Hello World!";
//  return 0;
//}
//`
	//	buildUserSolution(src, "golang")
//	buildUserSolution(src2, "cpp")
	//
	//	//input := strings.Join(args, " ")
	//	l := lexer.New(src)
	//	tok := l.NextToken()
	//	for tok.Type != token.EOF {
	//		fmt.Println(tok)
	//		if (tok.Type == token.IDENTIFIER){
	//			fmt.Println("i")
	//		}
	//		tok = l.NextToken()
	//	}
//}


// buildUserSolution returns error if user solution build failed
func buildUserSolution(solution, language string) (answer string, err error) {
	myErr := ""
	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, "", solution, 0)
	if err != nil {
		myErr = err.Error()
		fmt.Println(myErr)
		newErr := myErr[:strings.IndexByte(myErr, ':')]
		fmt.Println(newErr)
		answer += "Ошибка в "
		answer += newErr
		answer += " строке"

	} else {
		filePath, err := saveCodeFile("taskone", language, solution)
		if err != nil {
			fmt.Println(err)
		}

		//var buildCommands []string
		var compileCommands []string
		//cppBuildCommand     := "g++ -o"
		cppCompileCommand := "g++ program.cpp -o program"

		//golangBuildCommand  := "go build -o"
		golangCompileCommand := "go run"
		if language == "cpp" {
			//buildCommands = strings.Split(cppBuildCommand, " ")
			//compileCommands = strings.Split(cppCompileCommand, " ")
			//compileCommands = append(compileCommands, filePath)
			//compileCommands = strings.Split(cppCompileCommand, " ")
			cmd2 := exec.Command(cppCompileCommand)

			err = cmd2.Run()
			if err != nil {
				log.Println(err)
			}

			cmd3 := exec.Command("./program")
			out, err2 := cmd3.CombinedOutput()
			if err2 != nil {
				fmt.Println(err2)
			}
			fmt.Println(string(out))

		} else if language == "golang" {
			//buildCommands = strings.Split(golangBuildCommand, " ")
			compileCommands = strings.Split(golangCompileCommand, " ")
			compileCommands = append(compileCommands, filePath)
			cmd2 := exec.Command(compileCommands[0], compileCommands[1:]...)

			outq, err2 := cmd2.CombinedOutput()
			if err2 != nil {
				fmt.Println(err2)
			}
			fmt.Println("RESULT")
			fmt.Println(outq)
			fmt.Println(string(outq))
			answer = string(outq)
			//return string(out), err
		}

		// #nosec G204
		//buildCommands = append(buildCommands, binaryFilePath)
		//buildCommands = append(buildCommands, filePath)

		//fmt.Println(buildCommands)

		//cmd := exec.Command(buildCommands[0], buildCommands[1:]...)
		//
		//err = cmd.Run()
		//if err != nil {
		//	log.Println(err)
		//}
		//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		//cmd2 := exec.Command(compileCommands[0], compileCommands[1:]...)
		//
		//out, err2 := cmd2.CombinedOutput()
		//if err2 != nil{
		//	fmt.Println(err2)
		//}
		//fmt.Println("RESULT")
		//fmt.Println(out)
		//fmt.Println(string(out))

		//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

		//out, err2 := cmd.CombinedOutput()
		//if err2 != nil{
		//	fmt.Println(err2)
		//}
		//fmt.Println("RESULT")
		//fmt.Println(out)
		//fmt.Println(string(out))
		// build

		// remove unnecessary file
		//removeErr := os.Remove(filePath)
		//if removeErr != nil {
		//	log.Println(removeErr)
		//}
	}
	fmt.Println("!!!!!!!!!!!!!!")
	fmt.Println(answer)
	return answer, err
}

func saveCodeFile(fileName, language, code string) (string, error) {
	fmt.Println(code)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	var fileExtension string
	switch language {
	case "cpp":
		fileExtension = ".cpp"
	case "golang":
		fileExtension = ".go"
	default:
		return "", fmt.Errorf("ERROR: %s is not supported language by executor", language)
	}
	//binaryFilePath := "___go_build_app_test_callgraph /Users/vitalina_zhdanko/go/src/app/test/callgraph/src/solution"+ "/" + fileName


	filePath := fileName + fileExtension

	file, err := os.Create(filePath)
	if err != nil{                          // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1)                          // выходим из программы
	}
	defer file.Close()

	file.WriteString(code)

	return filePath, nil
}