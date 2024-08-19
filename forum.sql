/* Perhaps not the best approach for creating the DB file and apparently not compatible with sqlite CLI.
It seems using os.Create is a common way to do it; then, connect to the DB and execute the statements to
create the tables. */

create database if not exists forum;

create table users if not exists (
       id integer primary key,
       email varchar(30) not null unique,
       username varchar(20) not null unique,
       password varchar(40) not null
);

create table user_profile if not exists (
       id integer primary key,
       user_id integer foreign key references users(id) on delete cascade,
       content text
);

create table posts if not exists (
       id integer primary key,
       user_id integer foreign key references users(id),
       content text not null
);

create table comments if not exists (
       id integer primary key,
       post_id integer foreign key references posts(id),
       user_id integer foreign key references users(id),
       content text not null
);

create table post_reaction if not exists (
       id integer primary key,
       user_id integer foreign key references users(id),
       liked boolean
);

create table comment_reaction if not exists (
       id integer primary key,
       user_id integer foreign key references users(id),
       liked boolean
);

create table categories if not exists (
       id integer primary key,
       label varchar(15) unique not null
);
