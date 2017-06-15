RBBT TOPICS API Service
============

Installation
------------

```sh
make
make install
```

Configuration
-------------

Format:

```yml

listen: 3000 #Listen HTTP port

```

Starting service
----------------

```sh
bin/rbbt-service --c cfg/config.yml
```

Testing
-------

For the Convey launching (starts on 8180 port)
```sh
make convey
```

Launch tests
```sh
make test
```