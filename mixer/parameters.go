package mixer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var white = color.New(color.FgWhite).SprintFunc()

type Parameter struct {
	Type    string
	Prompt  string
	Label   string
	Default interface{}
	Choices []string
}

type Parameters map[string]Parameter

func (params Parameters) Prompt(ctx Context) Context {
	reader := bufio.NewReader(os.Stdin)

	for key, def := range params {
		fmt.Println(white(def.Prompt))
		// if def.Label {
		// 	fmt.Print(def.Label)
		// }
		fmt.Print("> ")
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		ctx[key] = strings.Trim(response, " \n")
	}
	return ctx
}
