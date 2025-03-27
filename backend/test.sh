#!/bin/bash
go build . &&
PORT=8700 JWT_KEY=114514 DB_USER=floodguard DB_NAME=floodguard DB_PASSWORD=Engagement SMTP_MAIL=floodguard@126.com SMTP_PASSWORD=NTbKihEAKrNLcuJm SMTP_SERVER=smtp.126.com SMTP_PORT=25 ./floodguard
