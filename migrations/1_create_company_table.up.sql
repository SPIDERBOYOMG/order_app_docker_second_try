create table company(
    company_id serial primary key,
    company_name varchar(255) not null
);

insert into company (company_name)
values
('Tech Innovations'), 
('Green Solutions'), 
('Global Enterprises'), 
('Local Crafts'), 
('Future Tech');