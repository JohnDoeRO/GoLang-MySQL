# GoLang-MySQL

Simple MySQL crud

I used a simple table:

CREATE TABLE `make` (
  `id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

ALTER TABLE `make`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);



In the golang file I used commands from SQL which are familiar to everyone.

You will need to get
go get github.com/go-sql-driver/mysql