package main

import "github.com/deanishe/awgo/keychain"

var Kc = keychain.New("alfred-" + Wf.BundleID())

type Result struct {
	Title    string
	Subtitle string
	Arg      string
}

func Run() {
	q := Wf.Args()[0]

	// search github for repositories e.g. "alfred in:name user:coheff"
	// or issues/pull requests e.g. "java is:pr is:merged"
	results := Search(q)
	for _, result := range results {
		Wf.
			NewItem(result.Title).
			Subtitle(result.Subtitle).
			Arg(result.Arg).
			Valid(true)
	}

	Wf.WarnEmpty("No results", "Try another search?")
	Wf.SendFeedback()
}
