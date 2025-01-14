# Examples

There are two examples in the `examples` folder. These examples are using docker services for hosts.

## JSON Example

```json
{
    "servers": [
        {
            "name": "uptimekuma.localhost",
            "locations": [
                {
                    "path": "/",
                    "to": {
                        "host": "uptimekuma",
                        "port": "3001",
                        "schema": "http"
                    }
                }
            ]
        },
        {
            "name": "nginx.localhost",
            "locations": [
                {
                    "path": "/",
                    "to": {
                        "host": "nginx",
                        "port": "80",
                        "schema": "http"
                    }
                }
            ]
        }
    ]
}
```

## YAML Example

```yaml
servers:
  - name: uptimekuma.localhost
    locations:
      - path: "/"
        to: 
          host: "uptimekuma"
          port: "3001"
          schema: "http"

  - name: nginx.localhost
    locations:
      - path: "/"
        to: 
          host: "nginx"
          port: "80"
          schema: "http"
```

### Example info

For both examples above and in `./examples`, `uptimekuma.localhost` will forward all URL paths to the `uptimekuma` docker service and `nginx.localhost` will forward all URL paths to the NGINX docker service.

### Running

A `compose.yml` and `Dockerfile` are provided to allowed running the examples locally. To do this run:

```bash
docker compose build
docker compose up -d
```

Once done, the URLs http://uptimekuma.localhost and http://nginx.localhost can be navigated to.

- By default the `examples/config.yaml` is used.
    - If you wish to use the JSON instead change `command: -c /examples/config.yaml --yaml` to: `command: -c /examples/config.json --json` for the `goproxy` service in the `compose.yml`. Then run the above `docker compose` commands (`docker compose build` only needed if not ran before).
