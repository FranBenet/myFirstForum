package dbaser

var insertStatements = map[string]string{
	"users": `insert into users (email, username, password) values
("jmadsen@uef.fi", "johnnyboy", "SbdfbWE345$"),
("hanrautio@suomi.fi", "hannehell", "$65/Hjbxz+"),
("robkay@kood.fi", "robkay", "P0_b;\@Gz"),
("runemaster@emp.fi", "runeberg", "Z_BY(#lp?=");`,
	"posts": `insert into posts (user_id, title, content) values
(1, "Happy new year!", "First! Let's get this forum going!"),
(4, "Our early roots", "I really recommend Murro's book on European early history."),
(2, "Modern man survival", "Jamie Oliver's latest book is a must for single men!"),
(3, "Stoicking our way", "A must read for every stoic out there, Meditations by Marcus Aurelius");`,
	"categories": `insert into categories (label) values
("fiction"),
("sci-fi"),
("novel"),
("history"),
("philosophy"),
("science"),
("cooking");`,
	"post_categs": `insert into post_categs values
(1, 2),
(2, 4),
(2, 5),
(3, 7),
(4, 5);`,
	"post_reactions": `insert into post_reactions values
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
	"comments": `insert into comments (post_id, user_id, content) values
(1, 2, "Second! Happy new year!"),
(1, 3, "All the best everyone! I have some ideas for the coming year!"),
(1, 4, "So nice to see the forum growing :) Cheers!"),
(2, 1, "Indeed a great read! I like how his views contrast with that of his supervisor."),
(2, 4, "He has really developed his ideas into a new framework."),
(3, 1, "The 15-min recipes are wonderful!"),
(3, 3, "It's amazing how we can prepare great meals with very little."),
(3, 4, "The ultimate test tonight: I'm cooking dinner for a date. Fingers crossed!"),
(4, 1, "All hail the emperor!");`,
}
