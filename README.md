[![Circle CI](https://circleci.com/gh/techvein/gozen.svg?style=svg&circle-token=c428c1d551315c708f8d2205042db799060c3732)](https://circleci.com/gh/techvein/gozen)

# What is this?

This is like a MVC web framework to create Go project that is a combination of the following.

* web framework: [gin-gonic](https://github.com/gin-gonic/gin)
* ORM: [dbr](https://github.com/gocraft/dbr)
* migration tool: [goose](https://bitbucket.org/liamstask/goose/)

We don't think this project is best practice, so we have some plans to improve this.

# Usage

If you use Docker Compose, to set up is easy.

```
% docker-compose up -d
% docker-compose exec web bin/bash
$ go run tools/init.go
$ bash ./build.sh
```

Check the response

```bash
% curl http://localhost:5000/api/user/profile
{"message": "ログインしてください。"}
```

## With Gogland

#### Setup GOPATH
1. Open Preferences(⌘,) -&gt; Go -&gt; GOPATH
2. Add `<your project path>/src/gozen` to Project GOPATH.


# Contribution

Contributions to this project are welcome.

