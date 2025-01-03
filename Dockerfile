# First stage: Build the application
FROM golang:1.23.2 AS build-stage

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o error-service .

# Second stage: Create the final image
FROM scratch

# Copy the built application from the first stage
COPY --from=build-stage /app/static /static
COPY --from=build-stage /app/error-service /error-service

# Set the entrypoint to the executable
ENTRYPOINT ["/error-service"]