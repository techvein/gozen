FROM mysql:5.7
COPY ./my.cnf /etc/mysql/conf.d/

RUN apt-get update \
    ## /bin/sh: 1: Syntax error: redirection unexpected
    && bash -c 'debconf-set-selections <<< "mysql-server mysql-server/root_password password rootpass"' \
    && bash -c 'debconf-set-selections <<< "mysql-server mysql-server/root_password_again password rootpass"' \
    && apt-get -y install mysql-client

## 外部からmysqlサーバにアクセスできるように
RUN sed -i -e"s/^bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/" /etc/mysql/mysql.conf.d/mysqld.cnf


# COPY ./setup.sh /tmp/setup.sh
# RUN /tmp/setup.sh
## ポート番号 3306 を外部に公開
EXPOSE 3306
