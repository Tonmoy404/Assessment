start:
	go run main.go

install_psql:
	sudo docker pull postgres
init_postgress: 
	sudo docker run --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

run_postgress:
	sudo docker start postgres-container

stop_postgress:
	sudo docker stop postgres-container

createdb:
	sudo docker exec -it postgres-container createdb -U postgres ecommerce

dropdb:
	sudo docker exec -it postgres-container dropdb -U postgres ecommerce

containers:
	sudo docker ps -a

remove:
	sudo docker rm 
