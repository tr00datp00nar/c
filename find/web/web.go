package web

import (
	"net/url"
	"strings"

	scriptish "github.com/ganbarodigital/go_scriptish"
)

func rawUrlEncode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

func searchBrave(query string) {
	q := "https://search.brave.com/search?q=" + query
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("xdg-open", q),
	)
	pipeline.Exec()
}

func searchGithub(query string) {
	q := "https://github.com/search?q=" + query
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("xdg-open", q),
	)
	pipeline.Exec()
}

func searchWikipedia(query string) {
	q := "https://en.wikipedia.org/wiki/" + query
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("xdg-open", q),
	)
	pipeline.Exec()
}

func searchStackOverflow(query string) {
	q := "https://stackoverflow.com/search?q=" + query
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("xdg-open", q),
	)
	pipeline.Exec()
}

func searchYoutube(query string) {
	q := "https://www.youtube.com/results?search_query=" + query
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("xdg-open", q),
	)
	pipeline.Exec()
}

func searchGoogle(query string) {
	q := "https://www.google.com/search?q=" + query
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("xdg-open", q),
	)
	pipeline.Exec()
}

func searchAmazon(query string) {
	q := "https://www.amazon.com/s?k=" + query
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("xdg-open", q),
	)
	pipeline.Exec()
}

func searchWolfram(query string) {
	q := "https://www.wolframalpha.com/input?i=" + query
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("xdg-open", q),
	)
	pipeline.Exec()
}
