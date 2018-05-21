FROM golang:onbuild
ENV PORT=80
EXPOSE 80/tcp

HEALTHCHECK --interval=5m --timeout=1s \
  CMD curl -f http://localhost/ || exit 1
