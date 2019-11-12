// models.article.go

package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	KeyWord string `json:"keyword"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	article{ID: 1, Title: "Gradle Blog", Content: "It is about using gradle", KeyWord: "Gradle"},
	article{ID: 2, Title: "GoLang Blog", Content: "It is about using go-gin", KeyWord: "Go-Gin"},
	article{ID: 3, Title: "Java Blog", Content: "It is about using java", KeyWord: "Java"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

// Fetch an article based on the ID supplied
func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

// Create a new article with the title and content provided
func createNewArticle(title, content string) (*article, error) {
	// Set the ID of a new article to one more than the number of articles
	a := article{ID: len(articleList) + 1, Title: title, Content: content}

	// Add the article to the list of articles
	articleList = append(articleList, a)

	return &a, nil
}
