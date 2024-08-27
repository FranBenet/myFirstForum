package dbaser

var insertStatements = []string{
	`insert into users (email, username, password) values
("jmadsen@uef.fi", "johnnyboy", "SbdfbWE345$"),
("hanrautio@suomi.fi", "hannehell", "$65/Hjbxz+"),
("robkay@kood.fi", "robkay", "P0_b;\@Gz"),
("runemaster@emp.fi", "runeberg", "Z_BY(#lp?=");`,
	`insert into posts (user_id, created, title, content) values
(1, "2023-12-31 19:09", "Happy new year!", "First! Let's get this forum going!"),
(4, "2024-01-01 14:32", "Our early roots", "I really recommend Murro's book on European early history."),
(2, "2024-01-03 20:07", "Modern man survival", "Jamie Oliver's latest book is a must for single men!"),
(3, "2024-01-11 17:57", "Stoicking our way", "A must read for every stoic out there, Meditations by Marcus Aurelius");`,
	`insert into categories (label) values
("fiction"),
("sci-fi"),
("novel"),
("history"),
("philosophy"),
("science"),
("cooking");`,
	`insert into post_categs values
(2, 1),
(4, 2),
(7, 3),
(5, 4);`,
	`insert into post_reactions values
(1, 1, 1),
(1, 2, 1),
(1, 3, 1),
(1, 4, 1),
(2, 1, 1),
(2, 2, 1),
(2, 3, 0),
(2, 4, 1),
(3, 1, 1),
(3, 3, 1),
(3, 4, 1),
(4, 1, 1),
(4, 4, 1),
(4, 3, 0);`,
}
