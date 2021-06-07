TAG=${TAG:-v0.1.2}
docker build -t miko/webgock .
docker tag miko/webgock miko/webgock:${TAG}

