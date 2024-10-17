ARG IMAGE=node:lts-alpine

FROM ${IMAGE} as builder
WORKDIR /app
COPY . .
RUN npm install --quiet --no-optional --no-fund --loglevel=error

FROM builder as prod-build
RUN npm run build
RUN npm prune --production

FROM ${IMAGE} as prod
COPY --from=prod-build /app/dist /app/dist
COPY --from=prod-build /app/node_modules /app/node_modules
COPY --from=prod-build /app/.env.production /app/dist/.env

WORKDIR /app/dist

CMD ["node", "./main.js"]