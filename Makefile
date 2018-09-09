db/start:
	docker-compose up -d

db/down:
	docker-compose down

mysql:
	mysql -u root -h localhost -P3306 --protocol tcp -ppassword