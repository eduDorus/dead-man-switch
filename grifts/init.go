package grifts

import (
	"github.com/edudorus/dead-man-switch/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
