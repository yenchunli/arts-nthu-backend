# NTHU Artscenter Server

## Packages Used

- `gonic/gin`
- `go-migration`

## Scripts

See Makefiles for more details.

### Create the Schema for Migration

```
migrate create -ext sql -dir db/migration -seq init_schema
```

### Create Database Docker Instance

```
docker run --name arts-db \
-p 5432:5432 \
-e POSTGRES_PASSWORD=Hello123@ \
-e POSTGRES_USER=hello \
-v /data/postgresql/data:/var/lib/postgresql/data \
-d postgres:13
```

