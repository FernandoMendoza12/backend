APP_NAME=broker-service
COMPOSE_FILE=docker-compose.yml

.PHONY: build up down restart logs clean prune

## Compila la imagen y levanta el stack
deploy: build up

## Construye las imágenes
build:
	docker-compose -f $(COMPOSE_FILE) build

## Levanta los contenedores
up:
	docker-compose -f $(COMPOSE_FILE) up -d

## Detiene los contenedores
down:
	docker-compose -f $(COMPOSE_FILE) down

## Reinicia el servicio
restart:
	$(MAKE) down
	$(MAKE) up

## Muestra los logs en vivo
logs:
	docker-compose -f $(COMPOSE_FILE) logs -f

## Borra contenedores, volúmenes y redes (cuidado)
clean:
	docker-compose -f $(COMPOSE_FILE) down -v --remove-orphans

## Elimina imágenes y cache no usadas
prune:
	docker system prune -af --volumes
