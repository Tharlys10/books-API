#!/bin/sh

export APP_PORT=3000
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=books
export DB_USER=developer
export DB_PASS=es12t13o2

rm -rv ./*.bin ; go build -v -o application.bin
./application.bin
