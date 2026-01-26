# workout api

intorduced things to the project durng development

go get https://github.com/jackc/pgx

go install github.com/pressly/goose/v3/cmd/goose@latest

go get -u github.com/stretchr/testify
no need to install
had to get for some reason also github.com/stretchr/testify/assert
didnt come together with it as of now

$ go get golang.org/x/crypto/bcrypt
goose -dir migrations postgres "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable" up


{
 "auth_token": {
  "token": "7GIBNNL5S7REOSISYDX6YEOVUK4QHSYYU3RO47OBNE4TTNRAASGQ",
  "expiry": "2026-01-27T20:11:37.829196+02:00"
 }
}
