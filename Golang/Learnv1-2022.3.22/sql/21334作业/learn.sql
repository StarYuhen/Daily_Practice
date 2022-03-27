# 创建学习表
create table learn(
    uid int  not null ,
    classphone int  not null ,
    # 成绩
    number int not null,
    primary key (uid,classphone)
);

