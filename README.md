# appoptics-demo
Demo app for python, node and golang with AppOptics support

# Run
```bash
docker-compose up
```

# Config
Place AppOptics key in .env file and start then
```bash
curl localhost:{3000,5000,8000}/{,redis,remote}
```
to start sending metrics.
