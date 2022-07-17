# Appointments

Appointments is a service for keep salon and user appointments.

## **How to run project?**

# What was used in the project?

* Golang
* Docker 
* Redis
* Mongo
* RabbitMQ

1. Create a file .env and setup env variables with name .env in folder build 
2. Create a file application.env and setup env variables with name in folder env
3. Execute:
~~~make
make docker
~~~

## **Setup envs before create container and runner application**
~~~env
API_HOST_PORT="0.0.0.0:8080"
API_GRACEFUL_WAIT_TIME="30s"

MONGO_HOST=
MONGO_USER=
MONGO_PASSWORD=
MONGO_DATABASE=
MONGO_COLLECTION=

REDIS_PASSWORD=
REDIS_HOST=

RABBIT_USER=
RABBIT_PASS=

SPLUNK_HOST=
SPLUNK_PASSWORD=
SPLUNK_TOKEN=
SPLUNK_SOURCE=
SPLUNK_SOURCETYPE=
SPLUNK_INDEX=
SPLUNK_PORT=


DD_ENV="development"
DD_SERVICE="scaffolding"
DD_WITH_PROFILER=false
~~~

## **Look at project progress on [kanban board](https://github.com/LeandroAlcantara-1997/beauty_salon_microsservices/projects/1)**