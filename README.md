# reto_tekcom


Req:

```
Crear Servicio llamado Multiplica
tendrá 2 parámetros ( numero1 , numero2 ) y devolverá los números multiplicados

Este deberá ser expuesto como un servicio Rest y gRPC ( pueden ser 2 módulos separados)
y deberá tener al menos 50% de test coverage.
```
##Introduction:
This project uses GO v.1.15.2

to execute the test : 

```console
    go test ./...
```

Overall coverage is > 50%, still more test are necessary to ensure the requirement

##Rest:
The rest service is started in main.go into root folder :

```console
 go run main.go
```

the following is a sample request:

```console
curl --location --request POST 'http://localhost:8080/v1/multiplication/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "a" : 100,
    "b" : 10
}'

```

#gRPC:
The gRPC service can be started using main.go into the directory _/gprc_, also 
there is a client wit a sample request into _/client_.

```console
 go run grpc/main.go
```

```console
 go run client/main.go
```
