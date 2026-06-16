package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("couldn't create request: %w", err)
	}

	req.Header.Set("User-Agent", "BootCrawler/1.0")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("couldn't fetch webpage: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("received error status code: %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("received a non-html content type: %s", contentType)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %w", err)
	}

	return string(body), nil
}

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	h1 := strings.TrimSpace(doc.Find("h1").First().Text())
	if h1 != "" {
		return h1
	}

	h2 := strings.TrimSpace(doc.Find("h2").First().Text())
	if h2 != "" {
		return h2
	}

	return ""
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	mainParagraph := strings.TrimSpace(doc.Find("main p").First().Text())
	if mainParagraph != "" {
		return mainParagraph
	}

	paragraph := strings.TrimSpace(doc.Find("p").First().Text())
	if paragraph != "" {
		return paragraph
	}

	return ""
}

func getURLsFromHTML(html string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	urls := []string{}

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		href = strings.TrimSpace(href)
		if href == "" {
			return
		}

		parsedURL, err := url.Parse(href)
		if err != nil {
			return
		}

		resolvedURL := baseURL.ResolveReference(parsedURL)
		urls = append(urls, resolvedURL.String())
	})

	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	images := []string{}

	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if !exists {
			return
		}

		src = strings.TrimSpace(src)
		if src == "" {
			return
		}

		parsedURL, err := url.Parse(src)
		if err != nil {
			return
		}

		resolvedURL := baseURL.ResolveReference(parsedURL)
		images = append(images, resolvedURL.String())
	})

	return images, nil
}

// func getHeadingsFromHTML(html string) ([]string, error) {
// 	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	headings := []string{}
//
// 	for i := 1; i <= 6; i++ {
// 		selector := "h" + strconv.Itoa(i)
//
// 		doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
// 			text := strings.TrimSpace(s.Text())
// 			if text != "" {
// 				headings = append(headings, text)
// 			}
// 		})
// 	}
//
// 	return headings, nil
// }
//
// func getParagraphsFromHTML(html string) ([]string, error) {
// 	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	paragraphs := []string{}
//
// 	doc.Find("p").Each(func(_ int, s *goquery.Selection) {
// 		text := strings.TrimSpace(s.Text())
// 		if text != "" {
// 			paragraphs = append(paragraphs, text)
// 		}
// 	})
//
// 	return paragraphs, nil
// }
