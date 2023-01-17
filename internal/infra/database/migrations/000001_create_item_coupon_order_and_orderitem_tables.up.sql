create schema if not exists ecommerce;

create table if not exists ecommerce.item
(
    id          serial primary key,
    category    text,
    description text,
    price       numeric,
    width       integer,
    height      integer,
    length      integer,
    weight      integer
);

create table if not exists ecommerce.coupon
(
    code        text,
    percentage  numeric,
    expire_date timestamp,
    primary key (code)
);

create table if not exists ecommerce.order
(
    id         serial,
    coupon     text,
    code       text,
    cpf        text,
    issue_date timestamp,
    freight    numeric,
    sequence   integer,
    total      numeric,
    primary key (id)
);

create table if not exists ecommerce.order_item
(
    id_order integer,
    id_item  integer,
    price    numeric,
    quantity integer,
    primary key (id_order, id_item)
);