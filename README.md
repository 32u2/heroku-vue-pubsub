# About

Quick starter with Go, Vue, Heroku and PubSub. Inspired by "decentralized Display fleet" use case brief.

# Stack

- Golang
- Vue
- Pubsub (Centrifuge)
- Heroku

# Deployment

Three ways to deploy:

1) Heroku git
2) GitHub - manual deployment
3) GitHub - automated deployment

# PubSub test

```
cd app
npm run build
cd ..
go run main.go
```

Test url: 127.0.0.1:5000/chat-test

# Command line refresher

Dependencies:
```
go mod vendor
go mod tidy
```

Heroku deployment:
```
git add .
git commit -am "my commit"
git push heroku main
```

Other Heroku:
```
heroku create <app-name>
heroku local
heroku logs --tail
heroku open
heroku apps:destroy my-app-name
go build -o bin/my-app-name -v .
$ git push heroku main
$ heroku open
```

## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
- [Heroku GitHub Integration](https://devcenter.heroku.com/articles/github-integration)
