FROM    envoyproxy/envoy-dev:6abcebc93f3d0616ba38ad9c5fec97dfc9c32645

RUN     apt update && apt install -y \
            luarocks

RUN     luarocks install luasocket

# Copy the Envoy configuration file into the image
COPY envoy.yaml /etc/envoy.yaml

# Create a directory for scripts
RUN mkdir -p /etc/scripts

# Copy the redis.lua script into the image
COPY redis.lua /etc/scripts/



# Specify the command to run on container startup
CMD ["envoy", "-c", "/etc/envoy.yaml"]
