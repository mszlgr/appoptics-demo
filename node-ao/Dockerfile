FROM node
COPY package-lock.json /app/
WORKDIR /app
RUN npm install

COPY . /app

ENTRYPOINT ["node"]
CMD ["server.js"]
