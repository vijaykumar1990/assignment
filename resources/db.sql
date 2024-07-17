create table mas_courses(id bigint, name varchar(50), PRIMARY KEY(id));
create table mas_students(email varchar(50), name varchar(50), PRIMARY KEY(email));

create table student_enrollment(email varchar(50), id bigint, signup_date date,status enum('active', 'inactive'), PRIMARY KEY(email, id), FOREIGN KEY (email) REFERENCES mas_students(email), 
FOREIGN KEY(id) REFERENCES mas_courses(id));
insert into courses.mas_courses(id,name) values(1,'CSE2345');
insert into courses.mas_students(email, name) values('vijay@gmail.com', 'ViJAY');
INSERT INTO `courses`.`student_enrollment`
(`email`,
`id`,
`signup_date`,
`status`)
VALUES
('vijay@gmail.com',
1,
'2024-07-13',
'active');