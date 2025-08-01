create table product_city (
    product_city_id serial primary key,
    product_id integer not null references product(product_id) on delete cascade,
    city_id integer not null references city(city_id) on delete cascade
);

INSERT INTO product_city (product_id, city_id)
VALUES
(1, 2),
(1, 4),
(2, 1),
(2, 5),
(3, 3),
(4, 1),
(4, 2),
(5, 4),
(5, 5);
