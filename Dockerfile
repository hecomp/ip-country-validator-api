FROM golang

RUN go get github.com/Masterminds/glide

RUN go get github.com/codegangsta/gin

WORKDIR /go/src/github.com/hecomp/ipCountryValidatorApi

ADD glide.yaml glide.yaml
ADD glide.lock glide.lock

RUN glide install

ADD  src src

CMD ["go", "run", "main.go"]