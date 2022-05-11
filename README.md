# visitors
Crawls the number of current visitors of Wellenbad Gleisdorf, Austria.

# query results
```
sqlite3 visitors.db
select DATETIME(created_at, 'unixepoch') as isodate, quantity from visits order by isodate DESC LIMIT 1;
```