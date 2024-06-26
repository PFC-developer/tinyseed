FROM --platform=linux/amd64 golang:1.22.1-alpine3.19 AS go-builder
RUN apk add ca-certificates build-base git cmake
#tzdata
WORKDIR /code/tinyseed
COPY . .
RUN go build -o bin/tenderseed

FROM --platform=linux/x86_64 alpine:3.19
RUN apk add --no-cache bash nano
RUN addgroup ts && adduser -G ts -D -h /ts ts
WORKDIR /ts
RUN mkdir /ts/.tenderseed
RUN chmod 777 /ts/.tenderseed
COPY --from=go-builder /code/tinyseed/bin/tenderseed /usr/local/bin/tenderseed
#COPY --from=go-builder /usr/share/zoneinfo/Asia/Almaty /etc/localtime
#RUN echo "Asia/Almaty" >  /etc/timezone
VOLUME /ts/.tenderseed
USER ts
CMD ["tenderseed"]
