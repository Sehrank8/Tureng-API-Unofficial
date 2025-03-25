package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

// Translation struct is any one translation of a word
type Translation struct {
	Category    string `json:"category"`
	Source      string `json:"source"`
	Translation string `json:"translation"`
	WordType    string `json:"word_type"`
}

// Response Ensures the response gives the word and count before the translations in the json
type Response struct {
	Word         string        `json:"word"`
	Count        int           `json:"count"`
	Translations []Translation `json:"translations"`
}

// Mapping of word type abbreviations to full names
var wordTypeMap = map[string]string{
	"i.":    "isim",
	"f.":    "fiil",
	"s.":    "sıfat",
	"zf.":   "zarf",
	"ünl.":  "ünlem",
	"expr.": "ifade",
}

// Extracts the full word type and cleans the source word
func extractWordType(source string) (string, string) {
	re := regexp.MustCompile(`(.*?)\s*([ifszünlexpr]+\.)$`)
	matches := re.FindStringSubmatch(source)

	if len(matches) == 3 {
		fullWordType, exists := wordTypeMap[matches[2]]
		if !exists {
			fullWordType = matches[2] // Fallback to original if not in map
		}
		return fullWordType, strings.TrimSpace(matches[1]) // Return full name + cleaned source
	}
	return "", source // If no match, return empty word type
}

// ScrapeTureng scrapes translations from the Tureng web page
func ScrapeTureng(word string) ([]Translation, error) {
	var translations []Translation

	c := colly.NewCollector()

	// Scrape the translations
	c.OnHTML("table#englishResultsTable tbody tr", func(e *colly.HTMLElement) {
		category := strings.TrimSpace(e.ChildText("td:nth-child(2)"))    // Column 2: Category
		source := strings.TrimSpace(e.ChildText("td:nth-child(3)"))      // Column 3: Turkish word
		translation := strings.TrimSpace(e.ChildText("td:nth-child(4)")) // Column 4: English translation

		// Extract word type (e.g., "book i." → "isim")
		wordType, source := extractWordType(source)

		// Append the translations to the translations Slice
		if source != "" && translation != "" {
			translations = append(translations, Translation{
				Category:    category,
				Source:      source,
				Translation: translation,
				WordType:    wordType,
			})
		}
	})
	// Visit the Tureng web page for the given word
	err := c.Visit(fmt.Sprintf("https://tureng.com/tr/turkce-ingilizce/%s", word))
	if err != nil {
		return nil, err
	}
	return translations, nil
}

// Handle API requests
func translateHandler(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "Missing 'word' query parameter", http.StatusBadRequest)
		return
	}

	//Translation fetcher
	translations, err := ScrapeTureng(word)
	if err != nil {
		http.Error(w, "Failed to fetch translations", http.StatusInternalServerError)
		return
	}

	// Set response header to JSON
	//w.Header().Set("Content-Type", "application/json")
	//
	//// Convert Go struct to JSON and send response !!!old code if you don't want the fixed indentation part
	//json.NewEncoder(w).Encode(map[string]interface{}{
	//Word:         word,
	//Count:        len(translations),
	//Translations: translations,
	//})

	//convert Translations struct to JSON and fix indentation for readability
	jsonData, err := json.MarshalIndent(Response{
		Word:         word,
		Count:        len(translations),
		Translations: translations,
	}, "", "  ")

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Write(jsonData)
}

func main() {
	// Set up the HTTP server
	http.HandleFunc("/translate", translateHandler)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
