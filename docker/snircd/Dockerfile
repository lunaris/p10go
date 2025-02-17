FROM ubuntu

RUN useradd -m snircd

RUN apt-get update && apt-get install -y wget gcc make python3 flex bison libpcre3-dev liblua5.1-0 liblua5.1-0-dev

WORKDIR /home/snircd
USER snircd

RUN wget --default-page=snircd.tar.gz https://api.github.com/repos/quakenet/snircd/tarball/ && tar -xf snircd.tar.gz && mv quakenet-snircd* snircd-src

WORKDIR /home/snircd/snircd-src

RUN ./configure -v --with-maxcon=10240 --enable-debug && make && make install

WORKDIR /home/snircd/bin
EXPOSE 6667
EXPOSE 4400

COPY ./ircd.conf /home/snircd/lib

CMD ["./ircd", "-n", "-x", "999"]
