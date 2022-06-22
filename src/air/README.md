# Air

Retrieve air quality data of city from https://api.waqi.info and cached with Redis. 


## Run the application 


```bash
export REDIS_SERVER_ADDRESS=<Redis Server>
export AQI_SERVER_URL=<AQI base URL>
export AQI_SERVER_TOKEN=<API Key>
export IP_STACK_SERVER_URL=<Lookup IP Serv>
export IP_STACK_SERVER_TOKEN=<API Key>

go run main.go

```