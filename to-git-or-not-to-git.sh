#! /bin/bash

curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"jamfrommars\"}}){id}}"}' | cut -d ":" -f4 | cut -d "}" -f1
