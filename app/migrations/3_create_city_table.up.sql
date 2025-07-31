create table city(
    city_id serial primary key,
    name varchar(255) not null
);

insert into city (name)
values 
('New York'), 
('Los Angeles'), 
('Chicago'), 
('Houston'), 
('Phoenix');