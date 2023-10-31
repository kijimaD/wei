# wei

Go template repository.

## install

```
$ go install github.com/kijimaD/wei@main
```

## image build

```
$ wei build
```

## record weight

```
$ wei -c config.yml rec 55.55
```

## docker run

```
$ docker run -v "$PWD/":/work -w /work --rm -it ghcr.io/kijimad/wei:latest
```
