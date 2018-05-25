FROM golang:1.9

RUN mkdir -p $GOPATH/src/amazon-downloader
WORKDIR $GOPATH/src/amazon-downloader/
COPY . .

RUN curl https://glide.sh/get | sh && glide install
RUN go-wrapper install amazon-downloader
RUN rm -rf ./*

ENV AMAZON_DOWNLOADER_HOST 0.0.0.0
ENV AMAZON_DOWNLOADER_PORT 8083

EXPOSE 8083

ENTRYPOINT amazon-downloader