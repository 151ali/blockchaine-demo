# blockchaine-demo
a blockchain demo written in GOlang

run the following command
```go 
 go run .
```

get the chain :
```bash
curl -X GET localhost:8080/chain
```

add a block :
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"sender":"person1","receiver":"prson2","amount":2}'
localhost:8080/add
```