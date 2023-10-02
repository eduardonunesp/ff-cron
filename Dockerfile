FROM golang as build

RUN mkdir -p /usr/build
WORKDIR /usr/build
COPY . .
RUN make build

FROM alpine

WORKDIR /usr/src/app
COPY --from=build /usr/build/bin/ffcron /ffcron

CMD [ "/ffcron" ]