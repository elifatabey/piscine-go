#! /bin/bash

curl -s https://learn.01founders.co/api/graphq1-engine/v1/graphq1 --data '{"query":"{user(where:{login:{_eg:\"atabeyelif\"}}){id}}"}'

