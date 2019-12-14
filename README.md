# Apache Flink Docker images

These images are part of the bigdata [docker image series](https://github.com/flokkr). All of the images use the same [base docker image](https://github.com/flokkr/docker-baseimage) which contains advanced configuration loading. 

It supports configuration based on environment variables (using specific naming convention), downloaded from consul and other plugins (for example to generate kerberos keystabs).

For more detailed instruction to configure the images see the [README](https://github.com/flokkr/docker-base/blob/master/README.md) in the flokkr/docker-base repository.

## Getting started

### Run

The easiest way to run a storm cluster is just use the included ```docker-compose.yaml``` file.

Checkout the repository and do a ```docker-compose up -d``` The Flink UI will be available at http://localhost:8081

You can adjust the settings in the compose-config file.




