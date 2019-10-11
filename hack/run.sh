#! bash
docker run --name redis_middleware_dev --network=host -d redis
# Start all main.go