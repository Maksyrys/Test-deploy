
DROP TABLE IF EXISTS reviews     CASCADE;
DROP TABLE IF EXISTS favorites   CASCADE;
DROP TABLE IF EXISTS cart_items  CASCADE;
DROP TABLE IF EXISTS order_items CASCADE;

DROP TABLE IF EXISTS orders      CASCADE;
DROP TABLE IF EXISTS books       CASCADE;

DROP TABLE IF EXISTS users       CASCADE;
DROP TABLE IF EXISTS categories  CASCADE;
DROP TABLE IF EXISTS authors     CASCADE;

DROP SEQUENCE IF EXISTS reviews_review_id_seq;
DROP SEQUENCE IF EXISTS cart_items_cart_item_id_seq;
DROP SEQUENCE IF EXISTS order_items_order_item_id_seq;
DROP SEQUENCE IF EXISTS orders_order_id_seq;
DROP SEQUENCE IF EXISTS users_user_id_seq;
DROP SEQUENCE IF EXISTS books_book_id_seq;
DROP SEQUENCE IF EXISTS categories_category_id_seq;
DROP SEQUENCE IF EXISTS authors_author_id_seq;
