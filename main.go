package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

type Annotation struct {
	Pk               string  `json:"pk"`
	Fragment         string  `json:"fragment"`
	Identifier       string  `json:"identifier"`
	IsPublic         bool    `json:"is_public"`
	Quote            string  `json:"quote"`
	Ranges           []Range `json:"ranges"`
	Text             string  `json:"text"`
	LastModifiedTime string  `json:"last_modified_time"`
	UserIdentifier   string  `json:"user_identifier"`
	ChapterTitle     string  `json:"chapter_title"`
	ChapterUrl       string  `json:"chapter_url"`
	CoverUrl         string  `json:"cover_url"`
	EpubIdentifier   string  `json:"epub_identifier"`
	EpubTitle        string  `json:"epub_title"`
}

type Range struct {
	End         string `json:"end"`
	Start       string `json:"start"`
	EndOffset   int    `json:"endOffset"`
	StartOffset int    `json:"startOffset"`
}

type Book struct {
	Title    string
	Chapters map[string][]QuoteAndText
}

type QuoteAndText struct {
	Quote string
	Text  string
}

func main() {
	// Open the JSON file
	file, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the JSON data into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the JSON data into a slice of Annotation structs
	var jsonStructure struct {
		Count    int          `json:"count"`
		Next     string       `json:"next"`
		Previous string       `json:"previous"`
		Results  []Annotation `json:"results"`
	}
	err = json.Unmarshal(data, &jsonStructure)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a map to store the annotations by book and chapter
	bookMap := make(map[string]Book)

	// Loop through the annotations and group them by book and chapter
	for _, annotation := range jsonStructure.Results {
		book := bookMap[annotation.EpubTitle]
		if book.Title == "" {
			book.Title = annotation.EpubTitle
			book.Chapters = make(map[string][]QuoteAndText)
		}
		chapter := book.Chapters[annotation.ChapterTitle]
		chapter = append(chapter, QuoteAndText{annotation.Quote, annotation.Text})
		book.Chapters[annotation.ChapterTitle] = chapter
		bookMap[annotation.EpubTitle] = book
	}

	// Convert the map to a slice of Book structs and sort it by title
	books := mapToBookSlice(bookMap)
	sort.Slice(books, func(i, j int) bool {
		return books[i].Title < books[j].Title
	})

	// Convert the sorted slice back to a map
	bookMap = make(map[string]Book)
	for _, book := range books {
		bookMap[book.Title] = book
	}

	for _, book := range bookMap {
		fmt.Printf("%s\n", book.Title)
		for chapterTitle, quotes := range book.Chapters {
			fmt.Printf("\t%s\n", chapterTitle)
			for _, quote := range quotes {
				fmt.Printf("\t\t%s\n", quote)
			}
		}
	}

	// Write the notes and highlights grouped by book and chapter to a file in markdown format
	err = writeNotesToMarkdown(bookMap, "notes.md")
	if err != nil {
		log.Fatal(err)
	}

}

// Helper function to convert a map to a slice of Book structs
func mapToBookSlice(bookMap map[string]Book) []Book {
	books := make([]Book, 0, len(bookMap))
	for _, book := range bookMap {
		books = append(books, book)
	}
	return books
}

func writeNotesToMarkdown(bookMap map[string]Book, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, book := range bookMap {
		_, err = file.WriteString(fmt.Sprintf("# %s\n", book.Title))
		if err != nil {
			return err
		}

		for chapterTitle, quotesAndTexts := range book.Chapters {
			_, err = file.WriteString(fmt.Sprintf("## %s\n", chapterTitle))
			_, err = file.WriteString(fmt.Sprintf("Amount of Notes: %d\n", len(quotesAndTexts)))
			if err != nil {
				return err
			}

			for _, quotesAndText := range quotesAndTexts {
				_, err = file.WriteString(fmt.Sprintf("- Quote: %s\n", quotesAndText.Quote))
				if quotesAndText.Text != "" {
					_, err = file.WriteString(fmt.Sprintf("- Comment: %s\n", quotesAndText.Text))
				}
				if err != nil {
					return err
				}
			}

			_, err = file.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
