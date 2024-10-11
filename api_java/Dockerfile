FROM eclipse-temurin:17-jdk-jammy as builder
WORKDIR /app
COPY . /app
RUN ./mvnw clean package

FROM openjdk:17
WORKDIR /app
COPY --from=builder /app/gateway/target/gateway-1.0.0.jar ./app.jar
ENTRYPOINT ["java", "-jar", "app.jar"]