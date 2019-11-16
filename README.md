# 🧪 Golang GraphQL experiment

```shell script
go get
go run .
```

## Example queries

**Get something**
```text
http://localhost:8080/graphql?query=%0A%7B%0A%20%20name%0A%7D&variables=%7B%0A%0A%7D
```

**Get something by ID**
```text
http://localhost:8080/graphql?query=%0A%7B%0A%20%20objectByID(id%3A%2042)%20%7B%0A%20%20%20%20id%0A%20%20%20%20name%0A%20%20%7D%0A%7D&variables=%7B%0A%0A%7D
```

