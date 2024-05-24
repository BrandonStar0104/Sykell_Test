package utils

import (
	"net/http"
	"net/url"
	"sykell/test/models"

	"github.com/PuerkitoBio/goquery"
)

func AnalyzeURL(pageURL string) models.ProcessResult {
	var result models.ProcessResult

	resp, err := http.Get(pageURL)
	if err != nil {
		return result
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return result
	}

	result.HTMLVersion = getHTMLVersion(doc)

	result.PageTitle = doc.Find("title").Text()
	result.HeadingsCount = countHeadings(doc)
	result.InternalLinks, result.ExternalLinks, result.InaccessibleLinks = countLinks(doc, pageURL)
	result.HasLoginForm = checkLoginForm(doc)

	return result
}

func getHTMLVersion(doc *goquery.Document) string {
	if goquery.NodeName(doc.Selection) == "html" {
		if doctype := doc.Nodes[0].FirstChild; doctype != nil && doctype.Type == 10 {
			if doctype.Data == "html" {
				return "HTML5"
			} else {
				return "Older HTML Version"
			}
		}
	}
	return "Unknown"
}

func countHeadings(doc *goquery.Document) map[string]int {
	headings := map[string]int{
		"h1": 0,
		"h2": 0,
		"h3": 0,
		"h4": 0,
		"h5": 0,
		"h6": 0,
	}
	for heading := range headings {
		headings[heading] = doc.Find(heading).Length()
	}
	return headings
}

func countLinks(doc *goquery.Document, pageURL string) (int, int, int) {
	internalLinks := 0
	externalLinks := 0
	inaccessibleLinks := 0

	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return internalLinks, externalLinks, inaccessibleLinks
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			linkURL, err := url.Parse(href)
			if err != nil {
				return
			}

			if linkURL.IsAbs() {
				externalLinks++
			} else {
				internalLinks++
				href = baseURL.ResolveReference(linkURL).String()
			}

			resp, err := http.Get(href)
			if err != nil || resp.StatusCode >= 400 {
				inaccessibleLinks++
			}
		}
	})

	return internalLinks, externalLinks, inaccessibleLinks
}

func checkLoginForm(doc *goquery.Document) bool {
	return doc.Find("input[type='password']").Length() > 0
}
