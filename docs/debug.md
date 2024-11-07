# Debugging

## Failed to connect to `localhost` port 8080

If when curling NuKV you get the following message:

```
curl 'localhost:8080/api/v1/_status'
curl: (7) Failed to connect to localhost port 8080 after 0 ms: Could not connect to server
```

There are two likely resolutions:

- Wait a little bit (the HTTP server provider does take a second or two to start up an dgrab `localhost:8080`
- Ensure that there are no *other* programs using the `8080` port (including previously launched providers -- `pgrep ghcr_io` should reveal those)
- Ensure there are no stale links left in lattice -- i.e. check `wash get links` for dubious/previously used links
