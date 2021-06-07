TAG=${TAG:-v0.1.1}
docker build -t miko/webgock .
docker tag miko/webgock miko/webgock:${TAG}

