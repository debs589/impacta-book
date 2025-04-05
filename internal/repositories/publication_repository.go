package repositories

import (
	"api/internal/models"
	"database/sql"
)

type DefaultPublicationRepository struct {
	db *sql.DB
}

func NewPublicationRepository(db *sql.DB) models.PublicationRepository {
	return &DefaultPublicationRepository{db}
}

func (r *DefaultPublicationRepository) CreatePublication(publication models.Publication) (int, error) {
	statement, err := r.db.Prepare("INSERT INTO publications(title, content, author_id) values(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *DefaultPublicationRepository) GetPublication(id int) (models.Publication, error) {
	publication := models.Publication{}

	row := r.db.QueryRow("SELECT p.*, u.nickName FROM publications p inner join users u on u.id = p.author_id WHERE p.id = ?", id)
	err := row.Scan(&publication.ID,
		&publication.Title,
		&publication.Content,
		&publication.AuthorID,
		&publication.Likes,
		&publication.CreatedAt,
		&publication.AuthorNick)

	if err != nil {
		return models.Publication{}, err
	}

	return publication, nil
}
