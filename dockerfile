FROM mifomm/golang-dev:1.14.7 as debug

WORKDIR /home/rh/go/src/work

COPY . . 

RUN sudo chown -R rh:rh /home/rh/go/src

ENV GOROOT=/usr/local/go
ENV GOPATH=/home/rh/go
ENV PATH=/home/rh/.nvm/versions/node/v14.8.0/bin:$PATH:$GOPATH/bin:$GOROOT/bin

RUN chmod +x build/debug.sh
RUN build/debug.sh

RUN chmod +x build/dlv.sh

ENTRYPOINT [ "build/dlv.sh" ]