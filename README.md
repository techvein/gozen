[![Circle CI](https://circleci.com/gh/techvein/gozen.svg?style=svg&circle-token=c428c1d551315c708f8d2205042db799060c3732)](https://circleci.com/gh/techvein/gozen)

# What is this?

This is like a MVC web framework to create Go project that is a combination of the following.

* web framework: [gin-gonic](https://github.com/gin-gonic/gin)
* ORM: [dbr](https://github.com/gocraft/dbr)
* migration tool: [goose](https://bitbucket.org/liamstask/goose/)

We don't think this project is best practice, so we have some plans to improve this.

# Usage

1.Create your project directory.

```bash
$ mkdir -p <your project>/src
```

2.Set GOPATH environment.

```bash
$ export GOPATH=<your project>
```

3.Git clone gozen

```bash
$ git clone https://github.com/techvein/gozen.git $GOPATH/src/gozen
```

4.After run mysql, setup mysql

```bash
$ cd $GOPATH/src/gozen
$ mysql -u root -prootpass < db/setup/mysql.sql
```

5.Run setup (install libraries and insert sample data into DB)

```bash
$ go run $GOPATH/src/gozen/tools/init.go
```

6.Run build.

```bash
$ cd $GOPATH/src/gozen
$ ./build.sh
```

7.Check the response

```bash
$ curl http://localhost:9000/api/users/1
{"Id":1,"Name":"田中"}
```

## With Intellij Idea

1. Open &lt;your project&gt; from File -&gt; Open... 
2. Setup GOPATH
    1. Open Preferences(⌘,) -&gt; Languages & Frameworks -&gt; Go -&gt; Go Libraries
    2. Add two paths(&lt;your project&gt; path and &lt;your project&gt;/src/gozen/vendor path) to Project libraries.
3. Run `$GOPATH/src/gozen/symlinkVendor.sh` to completion the libraries that installed by glide within Intellij Idea.   


## With Docker

TODO

## With vagrant

TODO

# Contribution

Contributions to this project are welcome.

