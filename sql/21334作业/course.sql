# 创建班级表
create table course
(
    courseid   int primary key not null,
    coursename text not null,
    coursels   text not null,
    coursemax  int  not null
);