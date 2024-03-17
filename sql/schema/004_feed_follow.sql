-- +goose Up
CREATE TABLE feed_follow (
    feed_id UUID NOT NULL REFERENCES feeds(id),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY (user_id, feed_id),
    created_at TIMESTAMP NOT NULL
);
-- +goose Down
DROP TABLE feed_follow;
