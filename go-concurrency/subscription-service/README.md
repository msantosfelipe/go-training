# subscription-service

### Run project:
- `docker compose up -d`
- `make start`
- Access `http://localhost:8090`

#### Access mailhog:
- Access `http://localhost:8025`

### Run project:
- `make stop`

### Run tests:
- `cd cmd/web`
- `go test -v .`
- `go test -race -v .` - test for race conditions
- `go test -coverprofile=coverage.out` - run tests and create a file with the coverprofile
- `go tool cover -html=coverage.out` - open the browser to visual see the coverage code

