DROP TABLE IF EXISTS books;
CREATE TABLE books (
    id              INT AUTO_INCREMENT PRIMARY KEY,
    title           VARCHAR(255),
    author          VARCHAR(255),
    first_name      VARCHAR(63),
    last_name       VARCHAR(63),
    genre           VARCHAR(127),
    series          VARCHAR(255),
    description     TEXT,
    publish_date    VARCHAR(16),
    publisher       VARCHAR(64),
    ean_isbn        VARCHAR(64),
    upc_isbn        VARCHAR(64),
    pages           SMALLINT UNSIGNED,
    ddc             VARCHAR(32),
    cover_style     VARCHAR(32),
    sprayed_edges   BOOLEAN,
    special_ed      BOOLEAN,
    first_ed        BOOLEAN,
    signed          BOOLEAN,
    location        VARCHAR(128)
);
