package dbaser

var insertUsers = `insert into users (email, username, password) values
("jmadsen@uef.fi", "johnnyboy", "SbdfbWE345$"),
("hanrautio@suomi.fi", "hannehell", "$65/Hjbxz+"),
("robkay@kood.fi", "robkay", "P0_b;\@Gz"),
("runemaster@emp.fi", "runeberg", "Z_BY(#lp?=");`

var insertPosts = `insert into posts (user_id, created, content) values
(1, "2023-12-31 19:09", "First! New Star Wars episode is awesome!"),
(4, "2024-01-01 14:32", "I really recommend Murro's book on European early history."),
(2, "2024-01-03 20:07", "Jamie Oliver's latest book is a must for single men!"),
(3, "2024-01-11 17:57", "A must read for every stoic out there, Meditations by Marcus Aurelius");`

var insertCategs = `insert into categories (name) values
("fiction"),
("sci-fi"),
("novel"),
("history"),
("philosophy"),
("science"),
("cooking");`

var postCategories = `insert into post_categs values
(2, 1),
(4, 2),
(7, 3),
(5, 4);`
