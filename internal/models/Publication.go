package models

type Publication struct {
	ID         int    `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorID   int    `json:"authorId,omitempty"`
	AuthorNick string `json:"authorNick,omitempty"`
	Likes      int    `json:"likes"`
	CreatedAt  string `json:"createdAt,omitempty"`
}

type PublicationService interface {
	CreatePublication(Publication) (int, error)
	GetPublication(id int) (Publication, error)
	GetPublications(id int) ([]Publication, error)
}

type PublicationRepository interface {
	CreatePublication(Publication) (int, error)
	GetPublication(id int) (Publication, error)
	GetPublications(id int) ([]Publication, error)
}
