FROM traefik:v2.6.3
ADD . /plugins-local/src/github.com/esenac/traefik-custom-router/
ADD ./traefik-config/traefik.yml /etc/traefik/traefik.yml
ADD ./traefik-config/dynamic-config.yml /etc/traefik/config/dynamic-config.yml