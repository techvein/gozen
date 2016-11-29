FROM golang:1.7.3

# 
ENV MY_PROJECT /myproject
RUN mkdir $MY_PROJECT


# gopath設定
RUN mkdir -p $MY_PROJECT/gopath
ENV GOPATH $MY_PROJECT/gopath

COPY ./ci-setting.sh $MY_PROJECT
RUN $MY_PROJECT/ci-setting.sh


ADD ./gopath $GOPATH

# goプログラム設定
RUN mkdir -p $MY_PROJECT/golang
RUN cp -r  /usr/local/go/ $MY_PROJECT/golang/


RUN mkdir -p $GOPATH/bin
# path設定
ENV PATH $GOPATH/bin:$MY_PROJECT/golang/go/bin:$PATH

# Install goose that is a migration tool
RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go get -u github.com/kardianos/govendor
RUN cd $GOPATH/src/gozen && govendor sync
# RUN go run $GOPATH/src/gozen/tools/setup.go


# サービスとしては必要ないが、調査したいときがあるため入れておきたい
RUN apt-get update && apt-get install -y net-tools
WORKDIR $GOPATH/src/gozen

# Unix syslog delivery error
# https://groups.google.com/a/codenvy.com/d/msg/codenvy/6K6SgvK09oQ/oPswTD5aCAAJ
RUN apt-get update -q &&  apt-get install -y rsyslog
ENTRYPOINT /usr/sbin/rsyslogd -n




# FROM buildpack-deps:jessie-scm

# RUN echo "---- golang start------------------------"




# # settings to use gozen
# RUN echo "---- gozen app start------------------------"

# RUN git clone https://github.com/techvein/gozen.git $GOPATH/src/gozen

# # Install goose that is a migration tool
# RUN go get bitbucket.org/liamstask/goose/cmd/goose

# RUN cd $GOPATH/src/gozen && glide up




# ## start mysql & setup database and user
# #TODO: 2016/02/04 00:15:58 dial tcp 127.0.0.1:3306: getsockopt: connection refused
# ADD ./mysql/setup_mysql.sh /tmp/setup_mysql.sh
# RUN /tmp/setup_mysql.sh && cd $GOPATH/src/gozen && goose up & sleep 6s


# ## ポート番号 3306 を外部に公開
# EXPOSE 3306



# TODO
#RUN cd $GOPATH/src/gozen && bash -c './symlinkVendor.sh`
#RUN cd $GOPATH/src/gozen && goose up

