![go_test](https://github.com/aeonva1ues/simple_server/actions/workflows/go.yml/badge.svg)

## Idea of this simple project is practice with docker, writing tests and run they in ci/cd
____
### Installation
Clone this repo
```
git clone https://github.com/aeonva1ues/simple_server
```
___
### Running
Use make to run tests
```
make test
```
or to run web server
```
make compose-up
```

Alternative way to run server without make
```
docker-compose up --build -d && docker-compose logs -f
```
___
### Endpoints
This practice not about writing web server, in this reason there is only one endpoint.

Send request with curl
```
curl localhost:8080/
```
If server is running correctly you should see response content `<body>testpage</body>`
___
### Stop server
To stop the server you can use make file too
```
make compose-down
```

If you want run this command in the same terminal where you started the server, you should to press CTRL + C before it