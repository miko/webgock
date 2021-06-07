TAG=${TAG:-v0.1.0}
docker build -t miko/webgock .
docker tag miko/webgock miko/webgock:${TAG}

