create table item (
	id serial primary key not null,
	name varchar(300) not null,
	price numeric(20, 2),
	inserted_at timestamp not null default current_timestamp
)

create table review (
	id serial primary key not null,
	name varchar not null,
	description varchar not null,
	inserted_at timestamp not null default current_timestamp
);