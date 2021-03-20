package main

import "errors"

type (
	article struct {
		ID int 			`json:"id"`
		Title string 	`json:"title"`
		Content string 	`json:"content"`
	}
)

var (
	articleList = []article {
		{ID:      1, Title:   "article 1", Content: "article 1 body"},
		{ID:      2, Title:   "article 2", Content: "article 2 body"},
	}
)

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}