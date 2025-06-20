
INSERT INTO categories (category_id, name) VALUES
                                               (2, 'Классическая литература'),
                                               (3, 'Фантастика / Антиутопия'),
                                               (4, 'Исторический роман'),
                                               (5, 'Драма')
    ON CONFLICT (category_id) DO NOTHING;


INSERT INTO authors (author_id, name, biography) VALUES
                                                     ( 1, 'Лев Толстой',           'Русский писатель, автор «Войны и мира», «Анны Карениной»…'),
                                                     ( 2, 'Фёдор Достоевский',     'Русский писатель, автор «Преступления и наказания»…'),
                                                     ( 3, 'Михаил Булгаков',       'Русский писатель, автор «Мастера и Маргариты»'),
                                                     ( 4, 'Александр Пушкин',      'Русский поэт и писатель, автор «Евгения Онегина»…'),
                                                     ( 5, 'Иван Тургенев',         'Русский писатель, автор «Отцов и детей»'),
                                                     ( 6, 'Николай Гоголь',        'Русский писатель, автор «Мёртвых душ»'),
                                                     ( 7, 'Иван Гончаров',         'Русский писатель, автор «Обломова»'),
                                                     ( 8, 'Михаил Лермонтов',      'Русский поэт и писатель, автор «Героя нашего времени»'),
                                                     ( 9, 'Михаил Шолохов',        'Русский писатель, автор «Тихого Дона»'),
                                                     (10, 'Борис Пастернак',       'Русский писатель, автор «Доктора Живаго»'),
                                                     (11, 'Александр Солженицын',  'Русский писатель, автор «Ракового корпуса»'),
                                                     (14, 'Ильф и Петров',         'Авторы «Золотого телёнка»'),
                                                     (15, 'Эрих Мария Ремарк',     'Немецкий писатель, автор «Трёх товарищей»'),
                                                     (16, 'Андрей Платонов',       'Русский писатель, автор «Чевенгура»'),
                                                     (18, 'Виктор Некто',          'Автор «Левиафана»'),
                                                     (19, 'Томас Харрис',          'Американский писатель, автор «Молчания ягнят»'),
                                                     (20, 'Маргарет Митчелл',      'Американская писательница, автор «Унесённых ветром»'),
                                                     (21, 'Джордж Оруэлл',         'Британский писатель, автор «1984»')
    ON CONFLICT (author_id) DO NOTHING;



INSERT INTO users (user_id, username, email, phone, password,
                   firstname, lastname, role)
VALUES
    (11, 'dr.podgornii@gmail.com', 'dr.podgornii@gmail.com', '+79287557296',
     '$2a$14$aqY7tGmTZpTsIQC7YSRi9.APoSJDhB9xEmQSw7ht63RkJKNRpAlsm',
     'Данил', 'Подгорный', 'admin')
    ON CONFLICT (user_id) DO NOTHING;


INSERT INTO books
(book_id, title, author_id, category_id, price,
 description, publish_date, image_url, detailed_description)
VALUES
    (31, 'Война и мир', 1, 2, 29.99,
     'Эпическая сага о Наполеоновских войнах', '1869-01-01',
     'https://www.litres.ru/book/lev-tolstoy/voyna-i-mir-kollekcionnoe-illustrirovannoe-izdanie-69495367/',
     'Подробное описание книги «Война и мир»'),

    (32, 'Преступление и наказание', 2, 2, 19.99,
     'Психологический роман о борьбе добра и зла', '1866-01-01',
     'https://www.mann-ivanov-ferber.ru/books/prestuplenie-i-nakazanie-young-adult/',
     'Подробное описание книги «Преступление и наказание»'),

    (33, 'Мастер и Маргарита', 3, 3, 24.50,
     'Фантастический роман с элементами сатиры', '1967-01-01',
     'https://eksmo.ru/book/master-i-margarita-ITD912912/',
     'Подробное описание книги «Мастер и Маргарита»'),

    (34, 'Анна Каренина', 1, 2, 22.50,
     'Роман о трагической любви', '1877-01-01',
     'https://azbooka.ru/books/anna-karenina-sexm',
     'Подробное описание книги «Анна Каренина»'),

    (35, 'Евгений Онегин', 4, 2, 15.00,
     'Роман в стихах о любви и судьбе', '1833-01-01',
     'https://azbooka.ru/books/evgeniy-onegin-r9la',
     'Подробное описание книги «Евгений Онегин»'),

    (36, 'Братья Карамазовы', 2, 2, 25.99,
     'Философский роман о семье и вере', '1880-01-01',
     'https://azbooka.ru/books/bratya-karamazovy-mjqk',
     'Подробное описание книги «Братья Карамазовы»'),

    (38, 'Отцы и дети', 5, 2, 17.99,
     'Роман о конфликте поколений', '1862-01-01',
     'https://eksmo.ru/book/ottsy-i-deti-ITD1278953/',
     'Подробное описание книги «Отцы и дети»'),

    (41, 'Герой нашего времени', 8, 2, 14.99,
     'Роман о судьбе человека в обществе', '1840-01-01',
     'https://www.ozon.ru/product/geroy-nashego-vremeni-oblozhka-s-ill-paracosm-nastya-kalyuzhnaya-lermontov-mihail-yurevich-1359953504/',
     'Подробное описание книги «Герой нашего времени»'),

    (60, '1984', 21, 3, 13.50,
     'Антиутопический роман о тоталитарном обществе', '1949-06-08',
     'https://azbooka.ru/books/master-i-margarita-r4as',
     'Подробное описание книги «1984»')
    ON CONFLICT (book_id) DO NOTHING;

SELECT setval('categories_category_id_seq', (SELECT MAX(category_id) FROM categories));
SELECT setval('authors_author_id_seq',     (SELECT MAX(author_id)    FROM authors));
SELECT setval('users_user_id_seq',         (SELECT MAX(user_id)      FROM users));
SELECT setval('books_book_id_seq',         (SELECT MAX(book_id)      FROM books));
