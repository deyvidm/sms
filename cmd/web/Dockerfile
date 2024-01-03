# Use the official Node.js 16 image.
# https://hub.docker.com/_/node
FROM node:16

# Create and change to the app directory.
WORKDIR /usr/src/app

# Copy application dependency manifests to the container image.
# A wildcard is used to ensure both package.json AND package-lock.json are copied.
# Copying this separately prevents re-running npm install on every code change.
COPY package*.json ./

# Install production dependencies.
RUN npm ci --only=production

# Copy local code to the container image.
COPY . ./

# Build the application
RUN npm run build

# Use a minimal image to run a production build of the app.
# This stage starts a new build stage and is not based on the Node.js image
FROM nginx:stable-alpine

# Copy the build artifacts from the Node.js image into the nginx image.
COPY --from=0 /usr/src/app/dist /usr/share/nginx/html

# Expose port 80 to the Docker host, so we can access it
# from the outside.
EXPOSE 80

# Run the web service on container startup.
CMD ["nginx", "-g", "daemon off;"]