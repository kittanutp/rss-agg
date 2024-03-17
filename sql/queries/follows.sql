-- name: Folows :one
    INSERT INTO feed_follow (user_id, feed_id, created_at)
    VALUES ($1, $2, $3)
    RETURNING *;

-- name: GetFollowFeeds :many
    SELECT feeds.id, feeds.user_id,  feeds.name, feeds.url, feeds.created_at, feeds.updated_at
    FROM feeds
    JOIN feed_follow ON feed_follow.feed_id = feeds.id
    WHERE feed_follow.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow WHERE feed_id = $1 and user_id = $2;

