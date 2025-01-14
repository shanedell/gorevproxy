# How to use gorevproxy

## Configuration

- The configuration for both JSON and YAML files is a list of servers under `servers:` (YAML) or `"servers":` (JSON). Each server has two attributes:
    - `name` which should be the hostname/URL wanting to be used
    - `locations`, a list of URL's path to forward. Each location has two attributes:
        - `path` specifies the URL path allowed to be used. If set to `""` or `/` all paths are allowed.
        - `to` specifies where the proxy is forwaring to. Each `to` has three attributes
            - `host` the host being forwarded to
            - `port` the port being forwarded to
            - `schema` the protocol to use for fowarding (`http` or `https`)

## Running

### Binary

Download the binary from https://github.com/shanedell/gorevproxy/releases

Run a service exposed to either localhost or docker service. This example will use an NGINX docker service, where the port is exposed to localhost.

Run NGINX:

```bash
docker run \
    --name gorevproxy-nginx \
    -d \
    -p 8080:80 \
    nginx:alpine
```

Create proxy config:

```bash
cat <<EOF > config.yml
servers:
  - name: "nginx.localhost"
    locations:
      - path: "/"
        to:
          host: "localhost"
          port: "8080"
          schema: "http"
EOF
```

Run `gorevproxy`:

```bash
./gorevproxy -c config.yml --yaml
```

Now you should be able to navigate to http://nginx.localhost

### Docker - Non-compose

This example focuses on running `docker run` commands instead of using `docker compose`.

First create a new docker network and run a service that's connect to that network.

Create network:

```bash
docker network create gorevproxy-net
```

Run NGINX:

```bash
docker run \
    --name gorevproxy-nginx \
    -d \
    --network gorevproxy-net \
    nginx:alpine
```

Create proxy config:

```bash
cat <<EOF > config.yml
servers:
  - name: "nginx.localhost"
    locations:
      - path: "/"
        to:
          host: "gorevproxy-nginx"
          port: "80"
          schema: "http"
EOF
```

Run `gorevproxy`

```bash
docker run \
    --name gorevproxy \
    -d \
    -v $PWD/config.yml:/config.yml \
    -p 80:80 \
    --network gorevproxy-net \
    ghcr.io/shanedell/gorevproxy:latest -c /config.yml --yaml
```

Now you should be able to navigate to http://nginx.localhost

### Docker - Compose

This example focuses on using `docker compose`.

Create `compose.yml` with the below content:

```docker-compose
services:
  nginx:
    image: nginx:alpine
    container_name: gorevproxy-nginx
    networks:
      - gorevproxy-net

  gorevproxy:
    image: ghcr.io/shanedell/gorevproxy:latest
    container_name: gorevproxy
    depends_on:
      - nginx
    command: -c /config.yml --yaml
    volumes:
      - ./config.yml:/config.yml
    ports:
      - "80:80"
    networks:
      - gorevproxy-net

networks:
  gorevproxy-net:
```

Create proxy config:

```bash
cat <<EOF > config.yml
servers:
  - name: "nginx.localhost"
    locations:
      - path: "/"
        to:
          host: "nginx"
          port: "80"
          schema: "http"
EOF
```

Run `docker compose`:

```bash
docker compose up -d
```

Now you should be able to navigate to http://nginx.localhost
