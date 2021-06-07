TAG=${TAG:-v0.1.4}
docker build --build-arg=TAG=$TAG -t miko/webgock .
docker tag miko/webgock miko/webgock:${TAG}

