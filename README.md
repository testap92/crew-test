## Crew Lambda Test

### Running locally:

```
$ cp .env.dist .env
$ docker-compose up --build
```

### Testing locally:

#### GET:
```
$ curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"httpMethod":"GET", "queryStringParameters":{"page":"1","limit":"1"}}' -H "Content-Type: application/json"
```

#### POST:
```
$ curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"httpMethod":"POST", "body":"{\"firstName\":\"hello-world\"}"}' -H "Content-Type: application/json"
```