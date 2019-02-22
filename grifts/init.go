package grifts

import (
  "github.com/gobuffalo/buffalo"
	"github.com/ricktimmis/motorhomeclub/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
