# Use an official Envoy Docker image as the base image
FROM envoyproxy/envoy:v1.21.0

# Copy the Envoy configuration file into the image
COPY envoy.yaml /etc/envoy.yaml

# Specify the command to run on container startup
CMD ["envoy", "-c", "/etc/envoy.yaml"]
