



-- +goose Up

CREATE TABLE post (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    by UUID,
    title VARCHAR(255),
    draft BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (by) REFERENCES users (id)
);

CREATE TABLE paragraph (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    post UUID,
    sub_title VARCHAR(255),
    text TEXT,
    FOREIGN KEY (post) REFERENCES post (id)
);

CREATE TABLE media (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    post UUID,
    link TEXT,
    sub_title VARCHAR(255),
    FOREIGN KEY (post) REFERENCES post (id)
);

-- +goose Down

DROP TABLE post;
DROP TABLE paragraph;
DROP TABLE media;




