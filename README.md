If you do not need docker swarm or k8s, but you really want to use docker secret, you have the alternative of use this project. 

You need assign two environment:
- "DIR_SECRET=/secret"
- "SECRET=MYSQL_PASSWORD_FILE,MYSQL_ROOT_PASSWORD_FILE"

The environment DIR_SECRET means where your secret will be saved, 
and the environment SECRET will be split into an array which contains your secret file name. 

you can see an example here:
* [docker-compose.yml](https://github.com/ssst0n3/docker_secret/blob/master/docker-compose.yml)

you can pull compiled images here: 
https://hub.docker.com/repository/docker/ssst0n3/docker_secret/general