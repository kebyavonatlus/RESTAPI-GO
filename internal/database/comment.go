package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kebyavonatlus/gorestapicourse/internal/comment"
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

func (database *Database) GetComment(
	ctx context.Context, uuid string,
) (comment.Comment, error) {

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
