# webservice-example

A simple web service example, running in Go, Java, JavaScript, PHP, Python and Ruby.

## How to Build

Install and configure Docker, then use `docker compose`:

```
docker compose up --build
```

Or

```
docker compose build
docker compose up -d
```

Open a browser and go to http://localhost/health or http://localhost/db-test

Each service is behind a proxy which is configured to load the next service (alphabetically) on each subsequent request, so open a URL in another tab or hit reload to use another version.
