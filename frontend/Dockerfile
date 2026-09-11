FROM node:22

WORKDIR /app

COPY package.json ./

RUN npm install

COPY . .

RUN rm -rf dist

RUN npm run build

EXPOSE 5173

ENTRYPOINT ["npm", "run"]