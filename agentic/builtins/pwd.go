package builtins

import (
	"strings"

	"github.com/qiangli/shell/agentic/shared"
)

type Pwd struct{}

func (b *Pwd) Match(line string) bool {
	return strings.HasPrefix(line, "pwd")
}

func (b *Pwd) Eval(g shared.IGosh, line string) error {
	path, err := g.GetWD()
	if err != nil {
		return err
	}

	g.WriteString(path)
	return nil
}
