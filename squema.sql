create database go_inventory;
use go_inventory;

create table USERS(
    id int auto_increment,
    email varchar(250) not null,
    name varchar(250) not null,
    password varchar(250) not null,
    primary key(id)
);

create table PRODUCTS(
    id int not null auto_increment,
    name varchar(250) not null,
    description varchar(250) not null,
    price float not null,
    created_by int not null,
    primary key(id),
    foreign key(created_by) references USERS(id)
);

create table ROLES(
    id int not null auto_increment,
    name varchar(250) not null,
    primary key(id)
);

create table USER_ROLES(
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key(id),
    foreign key(user_id) references USERS(id),
    foreign key(role_id) references ROLES(id)
);