FROM node:14-alpine

WORKDIR /app

COPY web/ .

RUN npm install

ENTRYPOINT [ "npm", "start" ]