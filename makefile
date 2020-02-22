################################################################################
## @Author:					Thomas Bouder <Tbouder>
## @Email:					Tbouder@protonmail.com
## @Date:					Sunday 05 January 2020 - 19:54:37
## @Filename:				makefile
##
## @Last modified by:		Tbouder
## @Last modified time:		Sunday 05 January 2020 - 19:55:03
################################################################################

.PHONY: members proxy pictures up proxy

all: members proxy pictures

init:
	@-echo "Cloning dependencies in ./src"
	@-mkdir src/
	@-git clone git@github.com:panghostlin/Pictures.git src/Pictures
	@-git clone git@github.com:panghostlin/Members.git src/Members
	@-git clone git@github.com:panghostlin/Proxy.git src/Proxy
	@-git clone git@github.com:panghostlin/Webapp.git src/Webapp
	@-go run install.go

update:
	@-(cd src/Pictures && git pull)
	@-(cd src/Members && git pull)
	@-(cd src/Proxy && git pull)
	@-(cd src/Webapp && git pull)

members:
	@-docker-compose stop panghostlin-members
	@-docker-compose build panghostlin-members
	@-docker-compose up -d --remove-orphans panghostlin-members
proxy:
	@-docker-compose stop panghostlin-proxy
	@-docker-compose build panghostlin-proxy
	@-docker-compose up -d --remove-orphans panghostlin-proxy
pictures:
	@-docker-compose stop panghostlin-pictures
	@-docker-compose rm panghostlin-pictures
	@-docker-compose build panghostlin-pictures
	@-docker-compose up -d --remove-orphans panghostlin-pictures
webapp:
	@-docker-compose stop panghostlin-webapp
	@-docker-compose build --no-cache panghostlin-webapp
	@-docker-compose up -d --remove-orphans --force-recreate panghostlin-webapp

re:
	docker-compose up -d --build --remove-orphans
down:
	docker-compose down
up:
	docker-compose up --build --remove-orphans
restart: down up

purge:
	@-rm -rf src/Pictures
	@-rm -rf src/Members
	@-rm -rf src/Proxy
	@-rm -rf src/Webapp
	@-rm -rf .data
	# @-rm -rf .pictures