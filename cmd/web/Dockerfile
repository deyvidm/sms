
FROM node:21-alpine as build

WORKDIR /usr/src/app

COPY package*.json ./

# Install production dependencies.
# TODO -- seperate dev and prod dependencies
RUN npm ci

# Copy local code to the container image.
COPY . ./

RUN npm run build

EXPOSE 3000

CMD ["node", "build"]