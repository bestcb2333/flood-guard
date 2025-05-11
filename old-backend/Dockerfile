FROM alpine:3.19.0
RUN apk update && apk add --no-cache libc6-compat tzdata
WORKDIR /app
COPY floodguard .
ENTRYPOINT [ "./floodguard" ]
