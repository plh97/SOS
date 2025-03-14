FROM node:22.13.1
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN apt-get update -y && apt-get install -y openssl && rm -rf /var/lib/apt/lists/*

COPY . /app
WORKDIR /app


RUN npm install -g corepack@latest
RUN corepack enable && corepack prepare pnpm@latest --activate

RUN pnpm -F db dev