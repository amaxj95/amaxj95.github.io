FROM alpine:3.16

RUN apk add --no-cache \
    openssh-server \
    nginx

EXPOSE 22 80

CMD ["/usr/sbin/nginx", "-g", "daemon off;"]