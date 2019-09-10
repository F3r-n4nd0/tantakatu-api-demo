FROM golang:1.12.9
WORKDIR  ./tantakatu-api/

ADD . .

RUN apt-get install -y curl \
  && curl -sL https://deb.nodesource.com/setup_10.x | bash - \
  && apt-get install -y nodejs \
  && curl -L https://www.npmjs.com/install.sh | sh

RUN npm install -g multi-file-swagger

RUN go get github.com/go-swagger/go-swagger

RUN go install github.com/go-swagger/go-swagger/cmd/swagger

RUN sh ./scripts/swagger_code_gen.sh

RUN go get ./... && go build -o tantakatu-api cmd/tantakatu-server/main.go


CMD sh ./scripts/run_services.sh
