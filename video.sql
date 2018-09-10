#create table users(
#id int unsigned primary key auto_increment,
#login_name varchar(64) unique key,
#pwd text
#)


#create table video_info(
#id varchar(64) primary key not null,
#author_id int unsigned,
#name text,
#display_ctime text,
#create_time datetime not null default current_timestamp) ;


create table comments(
id varchar(64) primary key not null,
video_id varchar(64),
author_id int unsigned,
content text,
time datetime not null default current_timestamp);

create table sessions(
session_id varchar(64) primary key not null,
ttl tinytext,
login_name varchar(64)
)

alter table 
