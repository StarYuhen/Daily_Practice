# 初始化架构 此文件只需要使用一次
CREATE SCHEMA database_test;

# 创建一个用户表
create table user
(
    # 自增，不为null
    id       int  not null auto_increment,
    name     text not null,
    password text not null,
    # 创建主键
    constraint primary key (id)
);

# 插入用户表数据
insert into user(name,password) values ('test','test'),
                                       ('test1','test1'),
                                       ('test2','test2'),
                                       ('test3','test3'),
                                       ('test4','test4');

# 创建一个价格表，同时将user的用户名作为外键
create table price
(
    id int not null auto_increment,
    price int not null,
    constraint primary key(id),
    # 同时关联外键
    foreign key (id) references user (id)
);

# 插入数据
insert into price(price) values (10),(20),(30),(40),(50);

