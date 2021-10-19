#!/bin/sh

export APP_PORT=3000
export JWT_SECRET=BOOK_API
export DB_HOST=
export DB_PORT=
export DB_NAME=
export DB_USER=
export DB_PASS=

rm -rv ./*.bin; go build -v -o application.bin
./application.bin
