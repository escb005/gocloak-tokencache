FROM quay.io/keycloak/keycloak:22.0

COPY . data/import
WORKDIR /opt/keycloak

ENV KC_HOSTNAME=localhost
ENV KEYCLOAK_USER=admin
ENV KEYCLOAK_PASSWORD=secret
ENV KEYCLOAK_ADMIN=admin
ENV KEYCLOAK_ADMIN_PASSWORD=secret

RUN /opt/keycloak/bin/kc.sh import --file /data/import/realm-export.json
ENTRYPOINT ["/opt/keycloak/bin/kc.sh"]