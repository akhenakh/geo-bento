# Geo-Bento

[Bento](https://warpstreamlabs.github.io/bento) plugins to transform geographic coordinates from a stream.

Bento is the swiss army of stream processing: Bento solves common data engineering tasks such as transformations, integrations, and multiplexing with declarative and unit testable configuration.

This repo contains multiple Bento plugins as Go modules, that you can build on demand (see `cmd/geo-bento`).

Note that the h3 plugin is using a [CGO free version](https://github.com/akhenakh/goh3).


## Get the country for a latitude and longitude

Use `country` with the following parameters: `latitude`, `longitude`.

An example `position.json`:

```js
{"id":42, "lat": 48.86, "lng": 2.34}
```

A `country.yaml` pipeline.

```yaml
input:
  file:
    paths: ["testdata/position.json"]
    codec: all-bytes

pipeline:
  threads: 1
  processors:
  - mapping: |
      #!blobl
      root = this
      root.country = country(this.lat, this.lng)

output:
  label: "out"
  stdout:
    codec: lines
```

Enrich the input with the timezone:

```sh
go build -o geo-0bento ./cmd/geo-bento
./geo-bento -c testdata/country.yaml
{"country":["France"],"id":42,"lat":48.86,"lng":2.34}
```

country module is using [coord2country](https://github.com/akhenakh/coord2country).

## Transform latitude and longitude into an Uber h3 cell

Use `h3` with the following parameters: `latitude`, `longitude`, `resolution`.

An example `position.json`:

```js
{"id":42, "lat": 48.86, "lng": 2.34}
```

A `h3.yaml` pipeline.

```yaml
input:
  file:
    paths: ["testdata/position.json"]
    codec: all-bytes

pipeline:
  threads: 1
  processors:
  - mapping: |
      #!blobl
      root = this
      root.h3 = h3(this.lat, this.lng, 5)

output:
  label: "out"
  stdout:
    codec: lines
```

Enrich the input with the h3 cell:

```sh
go build -o geo-bento ./cmd/geo-bento
./geo-bento -c testdata/h3.yaml
{"h3":"851fb467fffffff","id":42,"lat":48.86,"lng":2.34}
```

## Transform latitude and longitude into a Google s2 cell

Use `geos2` with the following parameters: `latitude`, `longitude`, `resolution`.

An example `position.json`:

```js
{"id":42, "lat": 48.86, "lng": 2.34}
```

A `s2.yaml` pipeline.

```yaml
input:
  file:
    paths: ["testdata/position.json"]
    codec: all-bytes

pipeline:
  threads: 1
  processors:
  - mapping: |
      #!blobl
      root = this
      root.s2 = s2_object(this.lat, this.lng, 15)

output:
  label: "out"
  stdout:
    codec: lines
```

Enrich the input with the s2 cell:

```sh
go build -o geo-bento ./cmd/geo-bento
./geo-bento -c testdata/s2.yaml
{"id":42,"lat":48.86,"lng":2.34,"s2":"2/033303031301002"}
```

## Get the timezone for a given latitude and longitude

Use `tz` with the following parameters: `latitude`, `longitude`.

An example `position.json`:

```js
{"id":42, "lat": 48.86, "lng": 2.34}
```

A `tz.yaml` pipeline.

```yaml
input:
  file:
    paths: ["testdata/position.json"]
    codec: all-bytes

pipeline:
  threads: 1
  processors:
  - mapping: |
      #!blobl
      root = this
      root.tz = tz(this.lat, this.lng)

output:
  label: "out"
  stdout:
    codec: lines
```

Enrich the input with the timezone:

```sh
go build -o geo-bento ./cmd/geo-bento
./geo-bento -c testdata/tz.yaml
{"tz":"Europe/Paris","id":42,"lat":48.86,"lng":2.34}
```

tz module is using [tzf](https://github.com/ringsaturn/tzf).

## Generate random position in a range (mainly for debug)

This is an input plugin that will generate random coordinates in your range.

```yaml
input:
  randpos:
    min_lat: 46.0
    max_lat: 48.0
    min_lng: 2.0
    max_lng: 2.3

output:
  label: "out"
  stdout:
    codec: lines
```

## Live Testing

Run this command and point your browser to http://localhost:4195/

```sh
./geo-bento blobl server --no-open --host 0.0.0.0 --input-file ./testdata/position.json -m testdata/s2_mapping.txt
```

## Docker

A pre built binary is also availale as a docker image:

```sh
 docker pull ghcr.io/akhenakh/geo-bento:latest
```

## TODO

- [ ] s2 shape index to perform PIP
- [ ] spatialite lookup to perform PIP
- [ ] random points in a rect
- [X] lat lng to h3
- [X] lat lng to s2
- [X] lat lng to tz
- [X] lat lng to country
