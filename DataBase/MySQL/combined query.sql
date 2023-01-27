select id,name from user where name like 'test%'
# union 表示组合查询，同时 union all 则表示返回所有匹配
union
select id,price from price where price>10;