# 创建课程表
create table class
(
    # 课程号主键
    classphone int           not null,
    classname  text          not null,
    classls    varchar(2000) not null check ( classls in ('公共课', '专业必修课', '选修课')),
    credit     int default 0 not null,
    startime   date          not null,
    tearche    text          not null
);

# 定义主键
alter table class
    add constraint class_pk
        primary key (classphone);