# minicrawler
Web Intelligence Class IT215 

Instructor : Eka Rahayu Setyaningsih, S.Kom., M.Kom. 

Students : Jeremy Kennedy - 223210693

# Docker Network
docker network create --driver=bridge --subnet=172.24.0.0/16 --ip-range=172.24.0.0/16 --gateway=172.24.1.1 docknet

# default postgres
docker run -d --name pgwisql -p 5450:5432 -e POSTGRES_DB=s2 -e POSTGRES_USER=s2usr -e POSTGRES_PASSWORD=s2usr -e PGDATA=${pwd}/docker/postgres/pgdata -v ${pwd}/docker/postgres:/var/lib/postgresql/data --network docknet --ip 172.24.1.110 postgres:16

# copy dump to container
docker cp docker/postgres/types.sql pgwisql:/home
docker cp docker/postgres/webintelligence.dump pgwisql:/home

# 
docker exec -it pgwisql /bin/bash

# restore
psql -h 172.24.1.110 -U s2usr -d s2 -f /home/types.sql
pg_restore -h 172.24.1.110 -U s2usr -d s2 /home/webintelligence.dump 


