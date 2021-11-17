# ha-task
This repository contains task related to ha assignment for Go Developer

## How to run 

Prerequisite : Docker

You can modify config/config.yml if you want to change sectorId or port

`$ docker build -t ha .`
`$ docker run -d -p 3003:3003 ha`
or
`$ go run main.go`

This will launch mysql database server, build Dockerfile and run app which will migrate and create DB.

- [Homepage](http://localhost:3003) :  http://localhost:3003  - This show realtime monitoring of the API Server using
Prometheus

- Swagger : http://127.0.0.1:3003/swagger/index.html
- ```
  curl -X 'POST' \
  'http://127.0.0.1:3003/api/v1/dns/location' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "vel": "20.0",
  "x": "123.12",
  "y": "456.56",
  "z": "789.89"
  }'
  ```

- To run test cases `$ go run ./routes`



## Additional questions :

- What instrumentation this service would need to ensure its observability and operational
transparency?

`I have added Prometheus that monitors number o connections, memory used and other parameters,
along with this we can add logging then create some alerts and notifications based on kind of failures`


- Why throttling is useful (if it is)? How would you implement it here?

`Throttling is important because it allows us to control traffic which might cause hardware failures 
or compromise from operating safe with security. TO implement this either control it from firewall if we have one
or using rate package from golang.org/x/time/rate we can utilize this and build custom throttling`


- What we have to change to make DNS be able to service several sectors at the same
time?

`We can make sectorID which I have configured in config.yml dynamic either make token based system or have them pass in payload`

- Our CEO wants to establish B2B integration with Mom's Friendly Robot Company by
allowing cargo ships of MomCorp to use DNS. The only issue is - MomCorp software
expects loc value in location field, but math stays the same. How would you
approach this? What’s would be your implementation strategy?

`Either design API based on token using that we can identify it's from MomCorp and send reponse object accordingly or ask client to send antoher param in input`

- Atlas Corp mathematicians made another breakthrough and now our navigation math is
even better and more accurate, so we started producing a new drone model, based on
new math. How would you enable scenario where DNS can serve both types of clients?

`For new drone model send some extra parameters so that we can separate both of them and accodingly use the equation`

- In general, how would you separate technical decision to deploy something from
business decision to release something? :

`First Analyse business requirement, see if any extra tables or fields required if yes change db schema accordingly,
then same for API and business logic see if we can use existing API or extend it, then write it and test test in dev after 
that push to staging and production, test properly in staging and once deployed to prod also, have CICD pipeline for this.`
