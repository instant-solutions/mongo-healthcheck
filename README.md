# MongoDB healthcheck
![Release](https://github.com/instant-solutions/mongo-healthcheck/actions/workflows/release.yml/badge.svg)
![Test](https://github.com/instant-solutions/mongo-healthcheck/actions/workflows/tests.yml/badge.svg)

Run performant and efficient MongoDB healthchecks without using the `mongosh` command. The `mongo` command line tool was marked as deprecated and removed since version 6. As a replacement `mongosh` was introduced, which is too [inefficient](https://jira.mongodb.org/browse/MONGOSH-1240) for healthchecks.

## Introduction

The command connects to the specified database and executes a `ping` command comparable to `mongosh --eval="db.adminCommand(\"ping\")"`.

```
$ ./mongo-healthcheck --help
Usage of ./mongo-healthcheck:
  -uri string
    	MongoDB connection string (default "mongodb://localhost:27017")
```

## Performance

### mongo-healthcheck

```
$ time mongo-healthcheck
Successfully pinged MongoDB.

real    0m0.013s
user    0m0.006s
sys     0m0.005s
```

### mongosh

```
$ time mongosh --eval="db.adminCommand(\"ping\")"

{ ok: 1 }

real    0m0.630s
user    0m0.615s
sys     0m0.078s
```

### mongo

```
$ time mongo --eval="db.adminCommand(\"ping\")"

{ "ok" : 1 }

real    0m0.065s
user    0m0.044s
sys     0m0.020s
```

## Contribution

Feel free to fork the project and send us a pull-request! :sunglasses:

Or consider sponsoring us, so we can continue to work on this project: [GitHub Sponsors](https://github.com/sponsors/instant-solutions) :star_struck:

## Author

Made with :heart: in Austria by [instant:solutions OG](https://instant-it.at)