#!/bin/bash
go build -o app . && \
PORT=8702 \
JWT_KEY=114514 \
DB_USER=patent \
DB_NAME=patent \
DB_PASS=patent \
DB_HOST=127.0.0.1 \
DB_PORT=3306 \
PATH_PATENT=/var/www/file/patent/ \
PATH_REPORT=/var/www/file/report/ \
DEEPSEEK_API=sk-f111f9380e2b42a5aed0bd206ed83cd6 \
./app
