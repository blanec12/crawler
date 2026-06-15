package main

import (
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{}
	}

	outgoingLinks, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		outgoingLinks = []string{}
	}

	imageURLs, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		imageURLs = []string{}
	}

	return PageData{
		URL:            baseURL.String(),
		Heading:        getHeadingFromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imageURLs,
	}
}
