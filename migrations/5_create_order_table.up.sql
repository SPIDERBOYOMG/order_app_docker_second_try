create table "order" (
    order_id serial primary key,
    price numeric(10, 2) not null,
    quantity integer not null,

    city_id integer not null references city(city_id),
    firm_id integer not null references firm(firm_id),
    company_id integer not null references company(company_id)
);

insert into "order" (price, quantity, city_id, firm_id, company_id) 
values
(100.00, 2, 1, 1, 1), 
(150.50, 3, 2, 2, 2), 
(200.75, 1, 3, 3, 3), 
(50.00, 5, 4, 4, 4), 
(300.00, 10, 5, 5, 5);