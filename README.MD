# Fruits API

Simple REST API in Go to manage fruits.

## Run Locally

```bash
go run .
```
app will be running on : http://localhost:8080/fruits

## Endpoints

- `GET /fruits` – list all fruits
- `GET /fruits/{id}` – get a fruit by ID
- `POST /fruits` – add a new fruit (use JSON like `{ "fruit": "banana", "color": "yellow" }`)
- in order to use cli following commands can be used to emulate POST request

```bash
 curl -X POST http://localhost:8080/fruits \
  -H "Content-Type: application/json" \
  -d '{"fruit": "banana", "color": "yellow"}
```  

## Build Docker

```bash
docker build -t fruits-api .
docker run -p 8080:8080 fruits-api
```

## Test
```bash
go test
```

