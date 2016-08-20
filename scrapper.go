package main

import (
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/yhat/scrape"
  "github.com/spf13/viper"
)

func AmazonScraper(amazonId string) (*Movie, jsonErr) {

  movie := &Movie{}

  resp, err := http.Get(viper.GetString("development.amazon_url") + amazonId)

  // amazon.de not responding
	if err != nil {
		return movie, jsonErr{Code: http.StatusInternalServerError, Text: "Site Unreachable"}
	}

	// item does not exist in amazon.de
	if resp.StatusCode == http.StatusNotFound {
		return movie, jsonErr{Code: http.StatusNotFound, Text: "Item Not Found"}
	}

	root, err := html.Parse(resp.Body)

  // page is not parsable or it doesn't have complete content
  if err != nil {
		return movie, jsonErr{Code: http.StatusPartialContent, Text: "Page Unparsable"}
	}

  // find title Node
	titleNode, _ := scrape.Find(root, scrape.ById("aiv-content-title"))
	movie.Title = scrape.Text(titleNode.FirstChild)

  // find releaseYear Node
	releaseYearNode, _ := scrape.Find(root, scrape.ByClass("release-year"))
	year, _ := strconv.Atoi(scrape.Text(releaseYearNode))
	movie.ReleaseYear = year

  // predicate actorsMatcher to find actorsNode
	actorsNode, _ := scrape.Find(root, actorsMatcher)
	movie.Actors = strings.Split(scrape.Text(actorsNode), ",")

  // predicate posterMatcher to find posterNode
	posterNode, _ := scrape.Find(root, posterMatcher)
	movie.Poster = scrape.Attr(posterNode, "src")

  // find similar movie Nodes
	movieNodes := scrape.FindAll(root, scrape.ByClass("downloadable_movie"))
	ids := make([]string, len(movieNodes))
	for i, movieNode := range movieNodes {
		ids[i] = scrape.Attr(movieNode, "data-asin")
	}
	movie.SimilarIDs = ids

	return movie, jsonErr{}
}

func actorsMatcher(n *html.Node) bool {
  if n.DataAtom == atom.Dd && n.Parent != nil &&
    n.PrevSibling != nil && n.PrevSibling.PrevSibling != nil {
    return scrape.Attr(n.Parent, "class") == "dv-meta-info size-small" &&
      scrape.Text(n.PrevSibling.PrevSibling) == "Darsteller:"
  }
  return false
}

func posterMatcher(n *html.Node) bool {
  if n.DataAtom == atom.Img && n.Parent != nil {
    return scrape.Attr(n.Parent, "class") == "dp-meta-icon-container"
  }
  return false
}
