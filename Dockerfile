############### BUILD STATE ######################
FROM golang:1.20.0 as BUILD
WORKDIR /app

ARG SqlNAME
ENV SqlNAME=$SqlNAME

# copy everything
COPY . .

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy files individually
COPY *.go ./
COPY api/*.go ./api/

# install 
RUN go install net/netip
RUN go build -ldflags="-s -x -X'main.apiVerion=v1'" -o /golang_guide

############### DEPLOY STATE #####################
FROM grc.io/distroless/base-debian10
WORKDIR /
COPY *.html ./
EXPOSE 80
USER nonroot:nonroot
ENTRYPOINT [ "/golang_guide" ]