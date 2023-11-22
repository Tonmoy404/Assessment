start:
	go run main.go

install_psql:
	sudo docker pull postgres

runpostgress: 
	sudo docker run --name tonmoy -e POSTGRES_PASSWORD=tonmoy404 -p 5432:5432 -d postgres

startpostgress:
	sudo docker start my-postgres

stoppostgress:
	sudo docker stop my-postgres

createdb:
	sudo docker exec -it my-postgres createdb -U postgres mydatabase

dropdb:
	sudo docker exec -it my-postgres dropdb -U postgres mydatabase



