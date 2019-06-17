# Anagramma

### Api
* Example request

Store data, method POST
```bash
curl localhost:8080/load -d '["foobar", "aabb", "baba", "boofar", "test"]'
```

Get data, method GET
```bash
curl 'localhost:8080/get?word=foobar'
```

### Build

darwin
```bash
make darwin
```

linux
```bash
make linux
```

linux and darwin
```bash
make build
```

### Run

```bash
./build/osx/anagramma -addr=localhost:8080
```