# docker_secret

## description
If you do not need docker swarm or k8s, but you really want to use docker secret, you have the alternative of use this project. 

You need assign two environment:
- "DIR_SECRET=/secret"
- "SECRET=MYSQL_PASSWORD_FILE,MYSQL_ROOT_PASSWORD_FILE,CERT_EXAMPLE"

The environment DIR_SECRET means where your secret will be saved, 
and the environment SECRET will be split into an array which contains your secret file name. 

If the environment SECRET contains a value starts from "CERT", this container will generate a pair of certificate and key.  
For example, if SECRET=CERT_EXAMPLE, there will be two files called EXAMPLE.CRT and EXAMPLE.key generated under DIR_SECRET.

The secrets will be generated randomly if secrets is not exists in DIR_SECRET otherwise not.

## quick-start
you can see an example here:
* [docker-compose.yml](https://github.com/ssst0n3/docker_secret/blob/master/docker-compose.yml)

you can pull compiled images here: 
* [ssst0n3/docker_secret](https://hub.docker.com/repository/docker/ssst0n3/docker_secret/general)

## for test
if you want to test your code, you can use environment DEVELOPMENT, we will copy all secrets to /tmp/secret/

you can see example here: 
* [docker-compose_test.yml](https://github.com/ssst0n3/docker_secret/blob/master/example/docker-compose_test.yml)
