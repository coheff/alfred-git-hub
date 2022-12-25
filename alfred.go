package main

func run() {
	q := wf.Args()[0]

	// search github
	// e.g. "alfred in:name user:coheff"
	results := search(q)
	for _, node := range results.Data.Search.Nodes {
		wf.
			NewItem(node.Name).
			Subtitle(node.URL).
			Arg(node.URL).
			Valid(true)
	}

	wf.WarnEmpty("No results", "Try another search?")
	wf.SendFeedback()
}
