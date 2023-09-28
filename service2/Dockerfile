FROM maven:3.8.5-openjdk-17 as build
COPY . /tmp
WORKDIR /tmp
RUN mvn clean install
FROM openjdk:17-alpine
RUN mkdir -p /app /logs
WORKDIR /app
COPY --from=build /tmp/target/*.jar /app/service2.jar
COPY --from=build /tmp/src/main/resources/* /app/
CMD ["/bin/sh","-c", "sleep 2; java -jar service2.jar --spring.config.location=/app/application.properties"]
