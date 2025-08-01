create table firm(
    firm_id serial primary key,
    firm_name varchar(255) not null
);

insert into firm (firm_name)
values 
('Tech Innovations'), 
('Green Solutions'), 
('Global Enterprises'), 
('Local Crafts'), 
('Future Tech');