# Use an official Envoy Docker image as the base image
FROM envoyproxy/envoy:v1.21.0

# Install Lua dependencies
RUN apt-get update && \
    apt-get install -y lua5.3 liblua5.3-dev luarocks && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Install lua-resty-redis module using luarocks
RUN luarocks install lua-resty-redis

# Set Lua package path to include lua-resty-redis module
ENV LUA_PATH="/usr/local/share/lua/5.3/?.lua;/usr/local/share/lua/5.3/?/init.lua;;"

# Copy the Envoy configuration file into the image
COPY envoy.yaml /etc/envoy.yaml

# Specify the command to run on container startup
CMD ["envoy", "-c", "/etc/envoy.yaml"]
