FROM golang:1.18 AS build

WORKDIR /opt/workflow/workctl
COPY . .
ARG GOPROXY="https://goproxy.cn"
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o workctl .

FROM alpine AS prod

RUN echo -e "http://mirrors.aliyun.com/alpine/v3.11/main\nhttp://mirrors.aliyun.com/alpine/v3.11/community" > /etc/apk/repositories \
    && apk add -U tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 

WORKDIR /opt/workflow/workctl

COPY --from=build /opt/workflow/workctl/workctl /opt/workflow/workctl/
COPY config /opt/workflow/workctl/
COPY template /opt/workflow/workctl/
COPY static /opt/workflow/workctl/
COPY docker/entrypoint.sh /opt/workflow/workctl/

RUN chmod 755 /opt/workflow/workctl/entrypoint.sh
RUN chmod 755 /opt/workflow/workctl/workctl

EXPOSE 8002
ENTRYPOINT [ "/opt/workflow/workctl/entrypoint.sh" ]