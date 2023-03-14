# go-todo

PostgreSQL 13
```
sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt update
sudo apt install postgresql-13 -y
systemctl enable postgresql
systemctl start postgresql
systemctl status postgresql
```

```

sudo -i -u postgres
psql

db
create table todo (
    todo_id serial primary key,
    title varchar(511) not null,
    category varchar(255),
    description varchar,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now() 
)

create function on_update() returns trigger as $$
    begin
        new.updated_at = now();
        return new;
    end;
$$ language plpgsql;


create trigger on_update_triger before update on todo
    for each row execute procedure on_update();

insert into todo (title, category, description) values ('title', 'cat', 'catcatcat');
update todo set title = 't', category = 'c', description = 'cat*3' where todo_id = 1;
delete from todo where todo_id = 1;
```