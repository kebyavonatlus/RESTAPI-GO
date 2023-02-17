package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kebyavonatlus/gorestapicourse/internal/comment"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(commentRow CommentRow) comment.Comment {
	return comment.Comment{
		ID:     commentRow.ID,
		Slug:   commentRow.Slug.String,
		Author: commentRow.Author.String,
		Body:   commentRow.Body.String,
	}
}

func (database *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var commentRow CommentRow
	row := database.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author FROM comments WHERE id = $1`,
		uuid,
	)
	err := row.Scan(&commentRow.ID, &commentRow.Slug, &commentRow.Body, &commentRow.Author)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching the comment by uuid: %w", err)
	}

	return convertCommentRowToComment(commentRow), nil
}

func (database *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	rows, err := database.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
		(id, slug, author, body)
		VALUES
		(:id, :slug, :author, :body)`,
		postRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return cmt, nil
}

func (database *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := database.Client.ExecContext(
		ctx,
		`DELETE FROM comments where id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to delete comment from database: %w", err)
	}

	return nil
}

func (database *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
	commentRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	rows, err := database.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		slug = :slug,
		author = :author,
		body = :body
		WHERE id = :id`,
		commentRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to update comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertCommentRowToComment(commentRow), nil
}
