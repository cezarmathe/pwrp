FROM golang

MAINTAINER Cezar Mathe <cezarmathe@gmail.com>

VOLUME /Volumes/pwrp_container /data

#copy configuration file
COPY .local/pwrp.toml /_pwrp/pwrp.toml

#download and build the binary
RUN git clone --depth=1 https://github.com/cezarmathe/pwrp.git /_pwrp/src
WORKDIR /_pwrp/src
RUN go build
RUN mv /_pwrp/src/pwrp /_pwrp/pwrp

ENTRYPOINT /_pwrp/pwrp
CMD ["--config /_pwrp/pwrp.toml"]