-- Users table
CREATE TABLE users
(
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username   VARCHAR(50)  NOT NULL UNIQUE,
    email      VARCHAR(255) NOT NULL UNIQUE,
    full_name  VARCHAR(100),
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    is_active  BOOLEAN          DEFAULT true
);

-- Posts table with foreign key to users
CREATE TABLE posts
(
    id           SERIAL PRIMARY KEY,
    user_id      uuid         NOT NULL,
    title        VARCHAR(255) NOT NULL,
    content      TEXT,
    status       VARCHAR(20) DEFAULT 'draft',
    published_at TIMESTAMP,
    view_count   INTEGER     DEFAULT 0,
    created_at   TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Comments table with foreign keys
CREATE TABLE comments
(
    id          SERIAL PRIMARY KEY,
    post_id     INTEGER NOT NULL,
    user_id     uuid    NOT NULL,
    content     TEXT    NOT NULL,
    is_approved BOOLEAN                  DEFAULT false,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    test_date   DATE                     DEFAULT CURRENT_DATE,
    test_time   TIME WITH TIME ZONE      DEFAULT CURRENT_TIME,
    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Tags table
CREATE TABLE tags
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    slug VARCHAR(50) NOT NULL UNIQUE
);

-- Post tags junction table
CREATE TABLE post_tags
(
    post_id INTEGER NOT NULL,
    tag_id  INTEGER NOT NULL,
    PRIMARY KEY (post_id, tag_id),
    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE
);