package model

import (
	"time"
)

// Manga represents a series for gql.
type Manga struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	OtherNames   []string   `json:"otherNames"`
	Description  string     `json:"description"`
	Demo         string     `json:"demo"`
	StartDate    time.Time  `json:"startDate"`
	EndDate      *time.Time `json:"endDate"`
	Genres       []string   `json:"genres"`
	ChapterCount int        `json:"chapterCount"`
	// would define these in separate model files
	// Chapters     []*Chapter  `json:"chapters"`
	// Mangakas     []*Mangaka  `json:"mangakas"`
	// Magazines    []*Magazine `json:"magazines"`
}
