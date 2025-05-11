# go-server
## project to check docker/k8s commands

## commands:
- `docker build -t sabyradinov/go-server:latest .` to create image with tag using main path;
- `docker build -f Dockerfile.dev .` to create image with specific file
- `docker run -p 8080:8080 sabyradinov/go-server` to create and exec new container with ports(int:ext) by tag name;
- `docker exec -it containerId sh` to exec container with rewrited inital command using instd and text beatify;
- `docker ps` to view list of runing containers;
- `docker stop containerId` to stop running container with id;

- `docker compose up -d` to run all containers using yml file, `-d` means background task
- `docker compose up --build` to rebuild and run containers using yml file
- `docker compose down` to stop all containers in yml file
- `restart policy` in docker compose, we've several policies to restart containers on failure, 'no', always, on-failure, unless-stopped;

- `docker run -p 8080:8080 -v ${pwd}:/app` - `-v` to make references from container to own folder