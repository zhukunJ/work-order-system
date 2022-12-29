FROM golang:1.18 AS build

WORKDIR /opt/workflow/ferry
COPY . .
ARG GOPROXY="https://goproxy.cn"
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ferry .

FROM alpine AS prod

RUN echo -e "http://mirrors.aliyun.com/alpine/v3.11/main\nhttp://mirrors.aliyun.com/alpine/v3.11/community" > /etc/apk/repositories \
    && apk add -U tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 

WORKDIR /opt/workflow/ferry

COPY --from=build /opt/workflow/ferry/ferry /opt/workflow/ferry/
COPY config /opt/workflow/ferry/
COPY template /opt/workflow/ferry/
COPY static /opt/workflow/ferry/
COPY docker/entrypoint.sh /opt/workflow/ferry/

RUN chmod 755 /opt/workflow/ferry/entrypoint.sh
RUN chmod 755 /opt/workflow/ferry/ferry

EXPOSE 8002
ENTRYPOINT [ "/opt/workflow/ferry/entrypoint.sh" ]