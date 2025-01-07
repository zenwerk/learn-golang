-- DBスキーマの定義を schema.sql を書く
CREATE TABLE authors (
    id   INTEGER PRIMARY KEY,
    name text    NOT NULL,
    bio  text
);