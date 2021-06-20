FROM node:14-alpine

WORKDIR /app

COPY web/ .

RUN npm ci

ENTRYPOINT [ "npm", "start" ]