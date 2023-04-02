FROM postgres:10.3

COPY ./resources/db/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]
