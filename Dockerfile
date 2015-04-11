#   docker run -p 80:80 -t antonikonovalov/avatars

FROM google/golang
ENV AV_HTTP 0.0.0.0:80

ADD . /gopath/src/github.com/antonikonovalov/avatars
WORKDIR /gopath/src/github.com/antonikonovalov/avatars

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv 7F0CEB10
RUN echo "deb http://repo.mongodb.org/apt/debian "$(lsb_release -sc)"/mongodb-org/3.0 main" | tee /etc/apt/sources.list.d/mongodb-org-3.0.list
RUN apt-get update
RUN apt-get install -y mongodb-org
RUN make

EXPOSE 80
ENTRYPOINT ["/gopath/src/github.com/antonikonovalov/avatars/avatars"]