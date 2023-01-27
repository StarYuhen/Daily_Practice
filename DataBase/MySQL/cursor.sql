# 创建储存过程
create procedure priceCursor()
begin
    # create cursor \\  cursor for :create sql
    declare users cursor for select id from user;
    # definition local variable
    declare total int default '1';
    # 打开游标
    open users;
    # 关闭游标
    close users;
end;

# use procedure cursor
call priceCursor();

# inquire database table ‘user’
select *
from user;

delete priceCursor();
