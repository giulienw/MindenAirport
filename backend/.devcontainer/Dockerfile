FROM --platform=linux/amd64 mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

SHELL ["/bin/bash", "-c"]
ENV TZ="Europe/Berlin"

RUN apt update && apt install -y libaio1 unzip
RUN mkdir -p /opt/oracle
RUN curl -L -O https://download.oracle.com/otn_software/linux/instantclient/2370000/instantclient-basic-linux.x64-23.7.0.25.01.zip
RUN unzip -o instantclient-basic-linux.x64-23.7.0.25.01 -d /opt/oracle
RUN rm instantclient-basic-linux.x64-23.7.0.25.01.zip
RUN curl -L -O https://download.oracle.com/otn_software/linux/instantclient/2370000/instantclient-sqlplus-linux.x64-23.7.0.25.01.zip
RUN unzip -o instantclient-sqlplus-linux.x64-23.7.0.25.01 -d /opt/oracle
RUN rm instantclient-sqlplus-linux.x64-23.7.0.25.01.zip
ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_23_7:$LD_LIBRARY_PATH
RUN mkdir -p /opt/oracle/instantclient_23_7/network/admin
RUN mkdir -p /oracle
COPY oracle/tnsnames.ora /opt/oracle/instantclient_23_7/network/admin/tnsnames.ora