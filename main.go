package main

import (
	aw "github.com/deanishe/awgo"
)

var Wf *aw.Workflow = aw.New()

func main() {
	Wf.Run(Run)
}
