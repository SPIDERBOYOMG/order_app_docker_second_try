create table product (
    product_id serial primary key,
    product_name varchar(255) not null,

    firm_id integer not null references firm(firm_id)
);

insert into product (product_name, firm_id)
values
('Smartphone X1', 1), 
('Eco-Friendly Gadget', 2), 
('Global Widget', 3), 
('Handcrafted Item', 4), 
('Next Gen Device', 5);