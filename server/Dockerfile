FROM node:18.12.1-alpine

# Upgrade packages
RUN apk update && \
    apk upgrade


WORKDIR /usr/src/app

# Install node dependencies
COPY package.json package-lock.json .
RUN npm install

# Copy application code
COPY index.js .
COPY src ./src

CMD ["npm", "run", "start"]
