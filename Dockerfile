FROM alpine:3.4
MAINTAINER Fabian Stegemann

LABEL org.label-schema.schema-version 1.0 \
      org.label-schema.name="cashier" \
      org.label-schema.description="cashier - A GitHub Authentication Gateway targeted at client only Web Applications" \
      org.label-schema.usage="/usr/share/doc/cashier/README.md" \
      org.label-schema.url="https://github.com/cashier" \
      org.label-schema.vcs-url="https://github.com/zetaron/cashier" \

EXPOSE 80

ENTRYPOINT ["/usr/bin/secret-wrapper", "/usr/bin/cashier"]

COPY secret-wrapper /usr/bin/secret-wrapper
COPY README.md /usr/share/doc/cashier/README.md
COPY config.yml.dist /etc/cashier/config.yml
COPY cashier /usr/bin/cashier

