# Use an official Envoy Docker image as the base image
FROM envoyproxy/envoy:v1.21.0

# Install Lua dependencies
RUN apt-get update && \
    apt-get install -y lua5.3 liblua5.3-dev && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Copy resty modules from OpenResty image
COPY --from=openresty/openresty:1.25.3.1-2-alpine-fat /usr/local/openresty/lualib/resty /usr/local/openresty/lualib/resty

# Copy the Lua filter script into the container
# COPY lua_filter.lua /etc/lua_filter.lua

# Copy the Envoy configuration file into the image
COPY envoy.yaml /etc/envoy.yaml

# Specify the command to run on container startup
CMD ["envoy", "-c", "/etc/envoy.yaml"]
