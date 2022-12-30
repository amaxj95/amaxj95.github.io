# ALPINE_DOCKER_POSTGRES IMAGE

## To build alpine docker and postgres image from docker file Type:

    #docker build . -t pg-alpine-dock

## To run alpine docker and postgres image from docker file Type:

    - docker run -d -v pgdoc_connect:/data/db -e POSTGRES_PASSWORD=pgadminpassword --name pgdoc_admin -p 5432:5432 pg-alpine-dock

    -- Prior to doing so, you may need to create a volume mount using 
    docker volume create pgdoc_connect
