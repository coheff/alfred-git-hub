package main

import (
	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow = aw.New()

func main() {
	wf.Run(Run)
}
