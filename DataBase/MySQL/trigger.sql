# create trigger
# trigger create need crud
# effect: create trigger in when insert user table,query id in price table
create trigger test
    after insert
    on user
    for each row select id
                 from price;

# effect: test trigger insert
insert into user(name, password)
values ('test_trigger', 'init');

# delete trigger
drop trigger test;