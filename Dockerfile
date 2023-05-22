FROM ubuntu:latest

WORKDIR /root

RUN apt-get update && \
    apt-get install -y sudo git curl wget vim tmux zsh fish && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir -p /root/.message

COPY ./bin/* /root/bin/

COPY .github/bash/welcome.sh /root/.message/welcome.sh

RUN chmod +x /root/.message/welcome.sh && \
    echo "/root/.message/welcome.sh" >> /root/.bashrc

CMD [ "/bin/bash" ]
