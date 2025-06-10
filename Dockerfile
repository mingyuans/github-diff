FROM golang:1.24-alpine3.21

WORKDIR /github-diff
COPY . /github-diff


RUN go mod download
RUN go mod tidy

RUN cd cmd/github-diff && go install

CMD ["github-diff"]