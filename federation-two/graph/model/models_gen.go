// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Content struct {
	ID string `json:"ID"`
}

func (Content) IsEntity() {}

type Review struct {
	ID      string     `json:"ID"`
	Rating  int        `json:"rating"`
	Content []*Content `json:"content"`
}

type ReviewInput struct {
	ID string `json:"id"`
}
