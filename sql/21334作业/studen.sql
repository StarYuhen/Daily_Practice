# 创建指定数据库架构
#create schema students21334;
# use students21334;
# 创建数据表 学生表

create table student
(
    # 自动增加uid--学号
    uid       int                       not null primary key,
    # 姓名
    name      text                      not null,
    # 使字段限制为男或女，默认女
    sex       varchar(2000) default '男' not null check (sex in ('男', '女')),
    # 籍贯
    gps       text                      not null,
    # 出生日期
    creattime date                      not null,
    # 备注
    ps        text,
    # 班级编号
    class     text                      not null
);

# 创建课程表
create table class
(
    # 课程号主键
    classphone int           not null primary key,
    classname  text          not null,
    classls    varchar(2000) not null check ( classls in ('公共课', '专业必修课', '选修课')),
    credit     int default 0 not null,
    startime   date          not null,
    tearche    text          not null
);


# 创建班级表
create table course
(
    courseid   int  not null primary key,
    coursename text not null,
    coursels   text not null,
    coursemax  int  not null
);

# 创建学习表
create table learn
(
    uids        int not null,
    classphones int not null,
    # 成绩
    number      int not null,
    primary key (uids, classphones)
);



/*
# 设置唯一
create unique index student_uindex_pk on student (uid);

# 创建自动增加
alter table student
    modify uid smallint auto_increment;
 */

# 定义外键
alter table learn
    add constraint student_learn_uids_fk
        foreign key (uids) references student (uid);



alter table learn
    add constraint learn_class_classphone_fk
        foreign key (classphones) references class (classphone);









