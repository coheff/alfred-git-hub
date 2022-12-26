package main

func Run() {
	q := wf.Args()[0]

	// search github for repositories e.g. "alfred in:name user:coheff"
	// or issues/pull requests e.g. "java is:pr is:merged"
	results := Search(q)
	for _, result := range results {
		wf.
			NewItem(result.Title).
			Subtitle(result.Subtitle).
			Arg(result.Arg).
			Valid(true)
	}

	wf.WarnEmpty("No results", "Try another search?")
	wf.SendFeedback()
}
