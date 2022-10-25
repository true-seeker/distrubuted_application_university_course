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
DELETE FROM students;
INSERT INTO students(name, emails, courses, birth_date, teachers, faculty, specialization)
VALUES ('1,Вася', 'a@yandex.ru|b@gmail.com', 'Математика|Русский язык|Английский', '01-01-2000',
        '1,Петр Петрович|2,Михаил Михайлович|3,Екатерина Екатериновна', 'Мехмат', 'КМБ'),
       ('2,Петя', 'c@yandex.ru', 'Математика|Философия', '02-02-2000',
        '1,Петр Петрович|4,Игорь Игоревич', 'ФСФ', 'Философия'),
       ('3,Лена', 'd@vk.com', 'Химия|География', '03-03-2000', '5,Кирилл Кириллович|6,Анна Анновна', 'Химический',
        'Прикладная химия');


SELECT * FROM students;


