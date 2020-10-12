If you do not need docker swarm or k8s, but you really want to use docker secret, you have the alternative of use this project. 

You need assign two environment:
- "DIR_SECRET=/secret"
- "SECRET=MYSQL_PASSWORD_FILE,MYSQL_ROOT_PASSWORD_FILE"

The environment DIR_SECRET means where your secret will be saved, 
and the environment SECRET will be split into an array which contains your secret file name. 

you can see a example here:
* [docker-compose.yml](docker-compose.yml)