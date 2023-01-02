#!/bin/bash
cd "$(dirname "$(realpath "$0")")";

echo "create database ${DBNAME};" > 99-initdb.sql
echo "create user ${DBUSER} password '${DBPASS}';" >> 99-initdb.sql
echo "grant all on database ${DBNAME} to ${DBUSER};" >> 99-initdb.sql

