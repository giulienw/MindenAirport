FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

SHELL ["/bin/bash", "-c"]

RUN apt update && apt install -y libaio1 unzip
RUN mkdir -p /opt/oracle
RUN curl -L -O https://download.oracle.com/otn_software/linux/instantclient/instantclient-basic-linux-arm64.zip
RUN unzip -o instantclient-basic-linux-arm64.zip -d /opt/oracle
RUN rm instantclient-basic-linux-arm64.zip
RUN curl -L -O https://download.oracle.com/otn_software/linux/instantclient/instantclient-sqlplus-linux-arm64.zip
RUN unzip -o instantclient-sqlplus-linux-arm64.zip -d /opt/oracle
RUN rm instantclient-sqlplus-linux-arm64.zip
RUN export LD_LIBRARY_PATH=/opt/oracle/instantclient_21_1:$LD_LIBRARY_PATH
RUN mkdir -p /opt/oracle/instantclient_21_1/network/admin
RUN mkdir -p /oracle
COPY oracle/tnsnames.ora /opt/oracle/instantclient_21_1/network/admin/tnsnames.ora