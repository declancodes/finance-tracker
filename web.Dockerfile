FROM node:14

WORKDIR /app

COPY web/ .

RUN npm install

ENTRYPOINT [ "npm", "start" ]