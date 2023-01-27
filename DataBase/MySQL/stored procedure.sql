# 使用存储过程

# 定义一个存储函数 接收三个参数，其中返回一个为变量
create procedure price(
    name_text text,
    price_int int,
    out price_id decimal(8, 2)
)
    # 定义过程函数
begin
    select max(price) into price_int from price;
#     select max(name) into name_text from user;
    select max(id) into price_id from price;
end;

# PROCEDURE price already exists 已经建立储存过程

call price('test', 1, @price_id);