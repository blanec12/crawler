package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	input := "<html><body><h1>Test Title</h1></body></html>"
	actual := getHeadingFromHTML(input)
	expected := "Test Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetHeadingFromHTMLH2Fallback(t *testing.T) {
	input := "<html><body><h2>Fallback Title</h2></body></html>"
	actual := getHeadingFromHTML(input)
	expected := "Fallback Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetHeadingFromHTMLNoHeading(t *testing.T) {
	input := "<html><body><p>No heading here.</p></body></html>"
	actual := getHeadingFromHTML(input)
	expected := ""

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLBasic(t *testing.T) {
	input := "<html><body><p>First paragraph.</p><p>Second paragraph.</p></body></html>"
	actual := getFirstParagraphFromHTML(input)
	expected := "First paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	input := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`

	actual := getFirstParagraphFromHTML(input)
	expected := "Main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLNoParagraph(t *testing.T) {
	input := "<html><body><h1>No paragraph here.</h1></body></html>"
	actual := getFirstParagraphFromHTML(input)
	expected := ""

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Fatalf("couldn't parse input URL: %v", err)
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLRelative(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body><a href="/blog">Blog</a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Fatalf("couldn't parse input URL: %v", err)
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com/blog"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLMultipleLinks(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body>
		<a href="/blog">Blog</a>
		<a href="/about">About</a>
		<a href="https://example.com/contact">Contact</a>
	</body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Fatalf("couldn't parse input URL: %v", err)
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"https://crawler-test.com/blog",
		"https://crawler-test.com/about",
		"https://example.com/contact",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLRelative(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body><img src="/logo.png" alt="Logo"></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Fatalf("couldn't parse input URL: %v", err)
	}

	actual, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body><img src="https://example.com/image.png" alt="Logo"></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Fatalf("couldn't parse input URL: %v", err)
	}

	actual, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://example.com/image.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLMissingSrc(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body>
		<img alt="No source">
		<img src="/logo.png" alt="Logo">
	</body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Fatalf("couldn't parse input URL: %v", err)
	}

	actual, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
