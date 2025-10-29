package builtins

import (
	"strings"

	"github.com/qiangli/shell/agentic/shared"
)

type Unset struct{}

func (b *Unset) Match(line string) bool {
	return strings.HasPrefix(line, "unset")
}

func (b *Unset) Eval(g shared.IGosh, line string) error {
	// TODO
	return nil
}
