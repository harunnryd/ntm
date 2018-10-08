## NTM (CQRS/es implementation in Go)

> Command query responsibility segregation (CQRS) applies the CQS principle by using separate Query and Command objects to retrieve and modify data, respectively.[2][3]
https://en.wikipedia.org/wiki/Command%E2%80%93query_separation

### How it works!
Hopefully you guys have installed **Docker** :)

```bash
docker-machine ls
docker-machine ip # Showing the IP eg: 192.168.99.100
vim docker-compose.yml # Open file
```

```yaml
  kafka:
    image: wurstmeister/kafka:0.10.2.1
    depends_on:
      - zookeeper
    environment:
      KAFKA_CREATE_TOPICS: "news-topic:1:3,statuses-topic:1:3,tags-topic:1:1:compact"
      KAFKA_ADVERTISED_HOST_NAME: 192.168.99.100 # The value must be change to match docker-machine ip
```

```bash
docker-compose run ntm ginkgo command commandApi # To run unit testing
docker-compose up --build # To run application
```
### Documentation
You can change the host to match docker-machine ip :)

------------


[Postman Documentation API Click here . . .](https://documenter.getpostman.com/view/5287012/RWgnZ1Hh "Postman Documentation API")

[![Postman](https://image.ibb.co/fEXozU/rest.png "Postman")

[![Terminal](https://image.ibb.co/iDmFeU/consumer.png "Terminal")

### ERD

[![ERD](https://image.ibb.co/fpAcR9/xasd.png "ERD")](https://image.ibb.co/fpAcR9/xasd.png "ERD")
