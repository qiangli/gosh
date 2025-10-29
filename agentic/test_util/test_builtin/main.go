package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/qiangli/shell/agentic/builtins/bkill"
	"github.com/qiangli/shell/agentic/sh"
)

func main() {
	g := sh.NewGosh()
	g.Init()
	defer g.Close()

	k := bkill.Kill{}
	line := strings.Join(os.Args, " ")

	err := k.Eval(g, line)

	if err != nil {
		fmt.Println("uncaught error:", err)
	}
}
