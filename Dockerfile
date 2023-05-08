FROM ubuntu:22.10

ARG binary=sender
ARG command=/sender
COPY $binary ./

CMD ["sh","-c","echo ${command}"]
