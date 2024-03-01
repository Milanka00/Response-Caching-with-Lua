# Use an official Envoy Docker image as the base image
FROM envoyproxy/envoy:v1.21.0

# Update package repositories and install necessary packages
# RUN apt-get update && \
#     apt-get install -y lua5.3 liblua5.3-dev luarocks && \
#     apt-get clean && \
#     rm -rf /var/lib/apt/lists/*

# Install Lua Redis library using luarocks
# RUN luarocks install lua-redis

# Copy the Envoy configuration file into the image
COPY envoy.yaml /etc/envoy.yaml

# Create a directory for scripts
RUN mkdir -p /etc/envoy/scripts

# Copy the redis.lua script into the image
COPY envoy_redis/redis.lua /etc/envoy/scripts/

# Specify the command to run on container startup
CMD ["envoy", "-c", "/etc/envoy.yaml"]
