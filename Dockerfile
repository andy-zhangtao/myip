FROM vikings/alpine
COPY bin/myip /myip
ENTRYPOINT ["/myip"]