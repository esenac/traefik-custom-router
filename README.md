# traefik-custom-router

A [Traefik](https://traefik.io) custom plugin to redirect users based on the Authorization header in the incoming request.

Reference: the official documentation for [developing Traefik custom plugins](https://doc.traefik.io/traefik-pilot/plugins/plugin-dev/).

## Usage

Set in the Dynamic config file (`traefik-config/dynamic-config.yml`) the URLs on which you want to redirect an anonymous user (`public` config)
and a logged one (`public` private).


# Build 

Execute 
```
docker build --tag traefik-custom-router .  
```
to build the Docker image.

# Run

Execute
```
docker run --rm -it -p 80:80 -p 8080:8080 --name traefik docker.io/library/traefik-custom-router
```
to see it in action.

### Build status
![Yaegi tests](https://github.com/esenac/traefik-custom-router/workflows/Go%20Matrix/badge.svg)
![Cross compilatiob](https://github.com/esenac/traefik-custom-router/workflows/Main/badge.svg)
