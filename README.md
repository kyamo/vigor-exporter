# vigor-exporter

A (crusty) DrayTek Vigor 130/165 (and others) exporter for prometheus.

This can be used to scrape your Vigor modem for some connection metrics.

## Build
```
go build .
```

## Usage Information
```
./vigor-exporter --help
  -host string
    	hostname/ip the Vigor is reachable on
  -password string
    	password to authenticate to the Vigor
  -username string
    	username to authenticate to the Vigor
```

The exporter is listening on `*:9103` and provides metrics at the `/metrics`
path.

## Docker
There is also a docker container available.
```
docker run -d -e HOST=$IP_OF_MODEM -e USERNAME=$USERNAME -e PASSWORD=$PASSWORD -p 9103:9103 kyamo/vigor_exporter
```

| Variable name | Description | Default value |
|---|---|---|
| HOST | Hostname/IP the Vigor is reachable on | 192.168.0.1 |
| USERNAME | Username to authenticate with the Vigor | admin |
| PASSWORD | Password to authenticate with the Vigor | changeme |
