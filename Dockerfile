FROM ubuntu:latest

MAINTAINER Deskmate neteasecloudmusicapi "190025254@qq.com"
RUN apt-key adv --keyserver pgp.mit.edu --recv-keys 573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62
RUN echo "deb http://nginx.org/packages/ubuntu/ trusty nginx" > /etc/apt/sources.list
RUN echo "deb-src http://nginx.org/packages/mainline/ubuntu/ trusty nginx" >> /etc/apt/sources.list
RUN echo deb http://mirrors.163.com/ubuntu/ trusty main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb http://mirrors.163.com/ubuntu/ trusty-security main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb http://mirrors.163.com/ubuntu/ trusty-updates main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb http://mirrors.163.com/ubuntu/ trusty-proposed main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb http://mirrors.163.com/ubuntu/ trusty-backports main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb-src http://mirrors.163.com/ubuntu/ trusty main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb-src http://mirrors.163.com/ubuntu/ trusty-security main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb-src http://mirrors.163.com/ubuntu/ trusty-updates main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb-src http://mirrors.163.com/ubuntu/ trusty-proposed main restricted universe multiverse >> /etc/apt/sources.list
RUN echo deb-src http://mirrors.163.com/ubuntu/ trusty-backports main restricted universe multiverse >> /etc/apt/sources.list

#ENV NGINX_VERSION 1.7.8-1~wheezy
#ENV DEBIAN_FRONTEND noninteractive


#RUN apt-get update && apt-get install nginx
RUN apt-get update
RUN apt-get install -y nginx
RUN apt-get install -y supervisor
RUN apt-get install -y ca-certificates
RUN apt-get install -y vim


COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf

RUN cd / && mkdir neteasecloudmusicapi && mkdir log && mkdir neteasecloudmusicapi/upload && touch log/server.log
COPY ./static /neteasecloudmusicapi/static
COPY ./swagger /neteasecloudmusicapi/swagger
COPY ./languages /neteasecloudmusicapi/languages
COPY ./neteasecloudmusicapi /neteasecloudmusicapi/neteasecloudmusicapi
COPY ./conf /neteasecloudmusicapi/conf

#COPY ./bin /neteasecloudmusicapi/bin

# forward request and error logs to docker log collector
#RUN ln -sf /dev/stdout /log/stdout.log
#RUN ln -sf /dev/stderr /log/stderr.log

VOLUME ["/var/cache/nginx","/log","/neteasecloudmusicapi/upload"]

EXPOSE 80 443 8080

CMD ["nginx", "-g", "daemon off;"]
CMD ["/neteasecloudmusicapi/neteasecloudmusicapi"]
CMD ["/usr/bin/supervisord"]
