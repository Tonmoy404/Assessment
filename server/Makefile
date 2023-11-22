start:
	go run main.go

install_psql:
	sudo docker pull postgres

runpostgress: 
	sudo docker run --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres

startpostgress:
	sudo docker start my-postgres

stoppostgress:
	sudo docker stop my-postgres

createdb:
	sudo docker exec -it my-postgres createdb -U postgres mydatabase

dropdb:
	sudo docker exec -it my-postgres dropdb -U postgres mydatabase



