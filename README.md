# Test lbc

Run a server using GIN Framework that computes a fizzbuzz request and retrieve the most called requests.

## Architecture

- Routes and handlers are defined in ```/handler/*```
- Middlewares and their implementations are defined in ```/middleware/*```
- The directory ```/utils/*``` contains utility functions
- The directory ```/binary-search-tree``` contains the binary search tree implementation, it actually only got AVL tree.
- The directory ```/cmd/*``` contains all the commands, it only got the run command.

## Endpoints

- ```/fizzbuzz```, the request paramaters are defined in ```handler/fizzbuzz/fizzbuzz.go```, return a array of string
- ```/fizzbuzz/stats```, it return an array of ```FizzbuzzStatsResponse``` defined in ```handler/fizzbuzz/stats.go```

## Using

### Run

```go run main.go```

### Test

```go test ./...```

Example requests: 

- ```http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz```
- ```http://localhost:8080/fizzbuzz/stats```

Example response: 

- for ```http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=25&str1=abc&str2=def```, response: ```["1","2","abc","4","def","abc","7","8","abc","def","11","abc","13","14","abcdef","16","17","abc","19","def","abc","22","23","abc","def"]```
- for ```http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=20&str1=fizz&str2=buzz```, response: ```["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz"]```

- for ```http://localhost:8080/fizzbuzz/stats```, response: ```[{"Int1":3,"Int2":5,"Limit":20,"Str1":"fizz","Str2":"buzz","Count":2},{"Int1":3,"Int2":5,"Limit":25,"Str1":"abc","Str2":"def","Count":2}]```

## Implementation detail

- The first idea that comes in my mind to count the fizzbuzz requests was to use a map. But no knowing the number of different possibles requests, use a map was a bad idea. So I have decided to use a Binary Search Tree, especially an AVL tree but I could have also use a Red Black Tree.
- For the same fizzbuzz request, I recalculate every time the result because the execution time is not heavy, otherwise I could have also save the result the first time the request is asked.
