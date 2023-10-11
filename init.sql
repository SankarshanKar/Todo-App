.open todos.db

create table todos (
    id integer primary key autoincrement not null,
    todo varchar(128),
    done boolean
);

insert into todos(todo, done) values('Hello!', false);