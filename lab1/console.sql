CREATE TABLE students
(
    id             integer primary key autoincrement,
    name           text,
    emails         text,
    courses        text,
    birth_date     text,
    teachers       text,
    faculty        text,
    specialization text
);

INSERT INTO students(name, emails, courses, birth_date, teachers, faculty, specialization)
VALUES ('Вася', 'a@yandex.ru|b@gmail.com', 'Математика|Русский язык|Английский', '01-01-2000',
        'Петр Петрович|Михаил Михайлович|Екатерина Екатериновна', 'Мехмат', 'КМБ'),
       ('Петя', 'c@yandex.ru', 'Математика|Философия', '02-02-2000',
        'Петр Петрович|Игорь Игоревич', 'ФСФ', 'Философия'),
       ('Лена', 'd@vk.com', 'Химия|География', '03-03-2000', 'Кирилл Кириллович|Анна Анновна', 'Химический',
        'Прикладная химия');


SELECT * FROM students;


