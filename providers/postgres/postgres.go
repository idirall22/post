package provider

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/idirall22/post/models"
	"github.com/lib/pq"
)

// PostgresProvider structure
type PostgresProvider struct {
	DB        *sql.DB
	TableName string
}

// New create a new post
func (p *PostgresProvider) New(ctx context.Context, content string, mediaURLs []string, userID, groupID int64) (*models.Post, error) {

	query := fmt.Sprintf(`
	    INSERT INTO %s (content, media_urls, user_id, group_id)
	    VALUES ($1, $2, $3, $4) RETURNING id, created_at
	    `, p.TableName)

	stmt, err := p.DB.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}
	post := &models.Post{}
	err = stmt.QueryRowContext(ctx, content, pq.Array(mediaURLs), userID, groupID).Scan(&post.ID, &post.CreatedAt)

	if err != nil {
		return nil, parseError(err)
	}
	post.Content = content
	post.UserID = userID
	post.GroupID = groupID

	return post, nil
}

// Get get a single post
func (p *PostgresProvider) Get(ctx context.Context, id, userID, groupID int64) (*models.Post, error) {

	// TODO: add a check by group and user
	query := fmt.Sprintf(`
        SELECT id, content, media_urls, user_id, group_id
        FROM %s
        WHERE id = %d AND group_id= %d AND deleted_at IS NULL
        `, p.TableName, id, groupID)

	stmt, err := p.DB.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	post := &models.Post{}
	err = stmt.QueryRowContext(ctx).
		Scan(
			&post.ID,
			&post.Content,
			(pq.Array)(&post.MediaURLs),
			&post.UserID,
			&post.GroupID,
		)

	if err != nil {
		return nil, err
	}
	return post, nil
}

// List return a list of posts using limit and offset
func (p *PostgresProvider) List(ctx context.Context, userID, groupID int64, limit, offset int) ([]*models.Post, error) {

	query := fmt.Sprintf(`
        SELECT id, content, media_urls, user_id, group_id FROM %s
        WHERE group_id = %d AND deleted_at IS NULL LIMIT %d OFFSET %d
        `, p.TableName, groupID, limit, offset*limit)

	stmt, err := p.DB.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	posts := []*models.Post{}

	for rows.Next() {
		post := &models.Post{}
		if err := rows.Scan(
			&post.ID,
			&post.Content,
			(pq.Array)(&post.MediaURLs),
			&post.UserID,
			&post.GroupID,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Update update a post
func (p *PostgresProvider) Update(ctx context.Context, content string, mediaURLs []string, userID, id int64) (*models.Post, error) {

	media, _ := pq.Array(mediaURLs).Value()

	query := fmt.Sprintf(`
        UPDATE %s SET content=$1, media_urls=$2
        WHERE id=$3 AND user_id=$4 AND deleted_at IS NULL
        RETURNING created_at
        `, p.TableName)

	stmt, err := p.DB.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	post := &models.Post{}
	err = stmt.QueryRowContext(ctx, content, media, id, userID).Scan(
		&post.CreatedAt,
	)

	post.ID = id
	post.Content = content
	post.MediaURLs = mediaURLs
	post.UserID = userID

	if err != nil {
		return nil, parseError(err)
	}
	return post, nil
}

// Delete delete a post
func (p *PostgresProvider) Delete(ctx context.Context, userID, id int64) error {

	stmt, err := p.DB.PrepareContext(ctx, "UPDATE posts SET deleted_at=$1 WHERE id=$2 AND user_id=$3")

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, time.Now(), id, userID)

	if err != nil {
		return parseError(err)
	}

	return nil
}
