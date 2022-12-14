# Лабораторная работа №2

___

На этапе проектирования системы необходимо выбрать предметную область и спроектировать нормализованную и
ненормализованную базу данных(БД). Ненормализованная БД должна храниться в формате SQLite, Microsoft Access, DBF или
другом формате настольной (файловой) БД и иметь одну таблицу. Нормализованная БД должна храниться в корпоративной СУБД –
PostgreSQL, Microsoft SQL Server, Oracle или MySQL – и иметь минимум пять таблиц.
Разработать распределенное приложение, в котором:

1. Сервис обмена данными должен выполнять прием данных в нормализованную БД. Непосредственно процесс нормализации может
   реализовывать либо программа экспорта, либо программа импорта (по выбору студента).
2. Данные передаются построчно при помощи сокетов и системы очередей сообщения (RabbitMQ, ActiveMQ, Yandex Message
   Queue) в зависимости от параметров запуска. Адреса подключения к другим компонентам распределённого приложения должны
   настраиваться при помощи конфигурационного файла.
3. Данные передаются по зашифрованному каналу связи. Шифрование сообщений может быть реализовано как самостоятельно, так
   и при помощи настройки SSL/TLS.
4. В случае самостоятельной реализации шифрования используется следующая схема:
    1. Данные перед передачей должны шифроваться при помощи ключа симметричного шифрования (AES, 256 бит).
    2. Ключ симметричного шифрования должен генерироваться программой экспорта и передаваться программе импорта для
       дешифрования данных.
    3. При этом ключ симметричного шифрования должен в свою очередь шифроваться при помощи ключа асимметричного
       шифрования (RSA, 2048 бит).
    4. Ключ асимметричного шифрования должен генерироваться программой импорта и программе экспорта должна передаваться
       открытая часть ключа.
       При сдаче лабораторной работы программы экспорта и импорта необходимо запустить на двух физически разных
       устройствах (рекомендовано – облачных).

Максимальное количество баллов, которые студент может получить за выполнение работы равно четырнадцати. Распределение
баллов представлено в следующей таблице:

| Требование к заданию                                                                    | Максимальное количество баллов    |
|-----------------------------------------------------------------------------------------|-----------------------------------|
| Распределённое приложение позволяет передавать информацию с помощью сокетов.            | 5                                 |
| Распределённое приложение позволяет передавать информацию с помощью очередей сообщений. | 5                                 |
| Данные передаются по зашифрованному каналу связи.                                       | 2 за каждый механизм коммуникации |