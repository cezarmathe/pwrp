FROM golang

MAINTAINER Cezar Mathe <cezarmathe@gmail.com>

COPY pwrp.toml /root/.config/pwrp.toml

#copy the source code
COPY . /_pwrp/src
WORKDIR /_pwrp/src

#build the binary
RUN go build -mod vendor
RUN mv /_pwrp/src/pwrp /_pwrp/pwrp

#runtime settings
ENTRYPOINT /_pwrp/pwrp