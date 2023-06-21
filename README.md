## OpenFeature お試し

```sh
docker run \
    -p 1031:1031 \
    -v $(pwd)/flag-config.yaml:/flag-config.yaml \
    -v $(pwd)/goff-proxy.yaml:/goff-proxy.yaml \
    thomaspoignant/go-feature-flag-relay-proxy:latest
```