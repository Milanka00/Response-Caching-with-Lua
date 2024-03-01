# Use an official Envoy Docker image as the base image
FROM envoyproxy/envoy:v1.21.0

# Copy the Envoy configuration file into the image
COPY envoy.yaml /etc/envoy.yaml

# Create a directory for scripts
RUN mkdir -p /etc/scripts

# Copy the redis.lua script into the image
COPY redis.lua /etc/scripts/
COPY socket.lua /etc/scripts/

# Specify the command to run on container startup
CMD ["envoy", "-c", "/etc/envoy.yaml"]
