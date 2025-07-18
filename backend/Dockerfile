# Use the devcontainer base setup for Oracle connectivity
FROM --platform=linux/amd64 mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

SHELL ["/bin/bash", "-c"]
ENV TZ="Europe/Berlin"

# Install Oracle Instant Client
RUN apt update && apt install -y libaio1 unzip && \
    mkdir -p /opt/oracle && \
    curl -L -O https://download.oracle.com/otn_software/linux/instantclient/2370000/instantclient-basic-linux.x64-23.7.0.25.01.zip && \
    unzip -o instantclient-basic-linux.x64-23.7.0.25.01 -d /opt/oracle && \
    rm instantclient-basic-linux.x64-23.7.0.25.01.zip && \
    curl -L -O https://download.oracle.com/otn_software/linux/instantclient/2370000/instantclient-sqlplus-linux.x64-23.7.0.25.01.zip && \
    unzip -o instantclient-sqlplus-linux.x64-23.7.0.25.01 -d /opt/oracle && \
    rm instantclient-sqlplus-linux.x64-23.7.0.25.01.zip && \
    mkdir -p /opt/oracle/instantclient_23_7/network/admin &&  \
    mkdir -p /oracle 

# Set Oracle environment variables
ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_23_7:$LD_LIBRARY_PATH
ENV PATH=/opt/oracle/instantclient_23_7:$PATH

# Copy tnsnames.ora for Oracle connection
COPY .devcontainer/oracle/tnsnames.ora /opt/oracle/instantclient_23_7/network/admin/tnsnames.ora

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
