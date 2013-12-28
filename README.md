## Apache benchmark

```
ab -n 5000 -c 100 http://localhost:3001/
```

## Starting Ruby server

```
bundle exec thin start -R app.ru -p 3001
```

## Starting Go server

```
GOMAXPROCS=4 go run app.go
```

## Starting Node.js server

```
node app.js
```
