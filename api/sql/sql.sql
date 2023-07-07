CREATE DATABASE IF NOT EXISTS devBook;
USE devBook;

CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    senha varchar(20) not null unique,
    criadoEm timestamp default current_timestamp()
)ENGINE=INNODB;
