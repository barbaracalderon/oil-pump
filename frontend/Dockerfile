FROM node:20-slim as node-base

WORKDIR /app

COPY . .

RUN npm install --force
RUN npm install @angular/cli -g 


EXPOSE 4200

CMD ["ng", "serve", "--host", "0.0.0.0", "--port", "4200"]