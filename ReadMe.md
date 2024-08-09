* Migration creation
* when we generate the first query, it generates db.go, models.go, querier.go along with <accounts>.sql.go
* issue
  * pgx/v5 => https://github.com/jackc/pgx/wiki/Getting-started-with-pgx
  * https://github.com/jackc/pgx/issues/743
* For pgx config using db connection pool see
  - https://github.com/techschool/simplebank
