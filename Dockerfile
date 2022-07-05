 FROM golang:1.18
#destination
 WORKDIR /app
#copy mod and sum for dependencies
 COPY go.mod ./
 COPY go.sum ./
#download dependencies
 
 RUN go mod download
#copy all other files 
 COPY . .

#[1/3] Writing a very simple forum in go [gorm, sqlite, bootstrap]\


 RUN go build -o /docker-vms-app
EXPOSE 8080
 CMD ["/docker-vms-app"]

