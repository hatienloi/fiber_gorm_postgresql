Compile and install packages and dependencies
```
go install
```

Add missing and remove unused modules
```
go mod tidy
```

Install Fiber cli
```
go install github.com/gofiber/cli/fiber@latest
```

RUN
```
docker-compose up -d
fiber dev
```