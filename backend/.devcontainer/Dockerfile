FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

FROM oraclelinux:8-slim

RUN  microdnf install oracle-instantclient-release-el8 && \
     microdnf install oracle-instantclient-basic oracle-instantclient-devel oracle-instantclient-sqlplus && \
     microdnf clean all