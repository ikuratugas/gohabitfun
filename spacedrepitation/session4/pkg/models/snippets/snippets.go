package snippets

import (
	"database/sql"
	"errors"

	"ikura.net/session4/pkg/models"
)

type SnippetMysql struct {
	DB *sql.DB
}

func (m *SnippetMysql) Latest() ([]*models.Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippetsss := []*models.Snippet{}
	for rows.Next() {
		s := &models.Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippetsss = append(snippetsss, s)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return snippetsss, nil
}

func (m *SnippetMysql) Insert(title, content, expires string) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	id, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, nil
	}

	result, err := id.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(result), nil
}

func (m *SnippetMysql) GetID(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`
	s := &models.Snippet{}

	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrSnippetNoRow
		} else {
			return nil, err
		}
	}

	return s, nil
}
