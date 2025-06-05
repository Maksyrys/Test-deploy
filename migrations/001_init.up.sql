
CREATE SEQUENCE authors_author_id_seq          AS integer START 1;
CREATE SEQUENCE categories_category_id_seq     AS integer START 1;
CREATE SEQUENCE books_book_id_seq              AS integer START 1;
CREATE SEQUENCE users_user_id_seq              AS integer START 1;
CREATE SEQUENCE orders_order_id_seq            AS integer START 1;
CREATE SEQUENCE order_items_order_item_id_seq  AS integer START 1;
CREATE SEQUENCE cart_items_cart_item_id_seq    AS integer START 1;
CREATE SEQUENCE reviews_review_id_seq          AS integer START 1;


CREATE TABLE authors (
                         author_id  integer      PRIMARY KEY      DEFAULT nextval('authors_author_id_seq'),
                         name       varchar(255) NOT NULL,
                         biography  text
);

CREATE TABLE categories (
                            category_id integer      PRIMARY KEY      DEFAULT nextval('categories_category_id_seq'),
                            name        varchar(100) NOT NULL UNIQUE
);

CREATE TABLE books (
                       book_id      integer      PRIMARY KEY      DEFAULT nextval('books_book_id_seq'),
                       title        varchar(255) NOT NULL,
                       author_id    integer REFERENCES authors(author_id)   ON DELETE SET NULL,
                       category_id  integer REFERENCES categories(category_id) ON DELETE SET NULL,
                       price        numeric(10,2) NOT NULL,
                       description  text,
                       publish_date date,
                       image_url    text,
                       detailed_description text
);

CREATE TABLE users (
                       user_id    integer       PRIMARY KEY      DEFAULT nextval('users_user_id_seq'),
                       username   varchar(100)  NOT NULL,
                       email      varchar(255)  NOT NULL UNIQUE,
                       phone      varchar(50),
                       password   varchar(255)  NOT NULL,
                       created_at timestamp     DEFAULT CURRENT_TIMESTAMP,
                       firstname  varchar(50),
                       lastname   varchar(50),
                       role       varchar(50)   NOT NULL DEFAULT 'user'
);

CREATE TABLE orders (
                        order_id   integer       PRIMARY KEY      DEFAULT nextval('orders_order_id_seq'),
                        user_id    integer       REFERENCES users(user_id),
                        order_date timestamp     DEFAULT CURRENT_TIMESTAMP,
                        total      numeric(10,2) NOT NULL,
                        status     varchar(50)   NOT NULL
);

CREATE TABLE order_items (
                             order_item_id integer PRIMARY KEY DEFAULT nextval('order_items_order_item_id_seq'),
                             order_id      integer REFERENCES orders(order_id),
                             book_id       integer REFERENCES books(book_id),
                             quantity      integer       NOT NULL,
                             price         numeric(10,2) NOT NULL
);

CREATE TABLE cart_items (
                            cart_item_id integer PRIMARY KEY DEFAULT nextval('cart_items_cart_item_id_seq'),
                            user_id      integer REFERENCES users(user_id) ON DELETE CASCADE,
                            book_id      integer REFERENCES books(book_id) ON DELETE CASCADE,
                            quantity     integer       NOT NULL DEFAULT 1,
                            added_at     timestamp     DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE favorites (
                           user_id  integer NOT NULL,
                           book_id  integer NOT NULL,
                           added_at timestamp DEFAULT CURRENT_TIMESTAMP,
                           PRIMARY KEY (user_id, book_id),
                           FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
                           FOREIGN KEY (book_id) REFERENCES books(book_id) ON DELETE CASCADE
);

CREATE TABLE reviews (
                         review_id  integer PRIMARY KEY DEFAULT nextval('reviews_review_id_seq'),
                         user_id    integer REFERENCES users(user_id)  ON DELETE CASCADE,
                         book_id    integer REFERENCES books(book_id)  ON DELETE CASCADE,
                         rating     integer CHECK (rating BETWEEN 1 AND 5),
                         comment    text,
                         created_at timestamp DEFAULT CURRENT_TIMESTAMP
);


