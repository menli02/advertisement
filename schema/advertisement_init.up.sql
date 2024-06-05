CREATE TABLE advertisements
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    title       varchar(255) not null,
    description varchar(255),
    price       decimal(10,2) not null default 0,
    createdTime timestamp,
    isActive        boolean      not null default false

);