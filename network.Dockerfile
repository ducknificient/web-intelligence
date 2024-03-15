# Use a base image
FROM ubuntu:latest

# Update apt and install necessary packages
RUN apt-get update && apt-get install -y \
    iproute2 \
    net-tools

# Create the network
RUN docker network create --driver=bridge --subnet=175.24.0.0/16 --ip-range=175.24.0.0/16 --gateway=175.24.1.1 imagenet

# Command to keep the container running
CMD tail -f /dev/null
