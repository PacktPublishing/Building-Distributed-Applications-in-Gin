FROM node:14.15.1
COPY package-lock.json .
COPY package.json .
RUN npm install
COPY . .
CMD npm start