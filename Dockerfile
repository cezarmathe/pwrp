FROM golang

MAINTAINER Cezar Mathe <cezarmathe@gmail.com>

VOLUME /Volumes/pwrp_container /data

COPY pwrp.toml /root/.config/pwrp.toml

#copy the source code
COPY . /_pwrp/src
WORKDIR /_pwrp/src

#run the tests
RUN go test -mod vendor

#build the binary
RUN go build -mod vendor
RUN mv /_pwrp/src/pwrp /_pwrp/pwrp

#runtime settings
ENV PWRP_DEBUG true

ENTRYPOINT /_pwrp/pwrp --verbose