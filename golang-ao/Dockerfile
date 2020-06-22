FROM golang
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . /go/src/golang-ao
WORKDIR /go/src/golang-ao
ENV GOPATH /go/
RUN dep ensure
RUN make build

CMD './server'
