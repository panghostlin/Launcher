################################################################################
## @Author:					Thomas Bouder <Tbouder>
## @Email:					Tbouder@protonmail.com
## @Date:					Sunday 05 January 2020 - 19:54:37
## @Filename:				makefile
##
## @Last modified by:		Tbouder
## @Last modified time:		Sunday 05 January 2020 - 19:55:03
################################################################################

.PHONY: members proxy pictures keys up proxy

all: keys members proxy pictures

init:
	@-echo "Cloning dependencies in ./src"
	@-mkdir src/
	@-git clone git@github.com:panghostlin/Pictures.git src/Pictures
	@-git clone git@github.com:panghostlin/Members.git src/Members
	@-git clone git@github.com:panghostlin/Keys.git src/Keys
	@-git clone git@github.com:panghostlin/Proxy.git src/Proxy
	@-git clone git@github.com:panghostlin/Webapp.git src/Webapp
	@-go run install.go

update:
	@-(cd src/Pictures && git pull)
	@-(cd src/Members && git pull)
	@-(cd src/Keys && git pull)
	@-(cd src/Proxy && git pull)
	@-(cd src/Webapp && git pull)

keys:
	@-docker-compose stop panghostlin-keys
	@-docker-compose build panghostlin-keys
	@-docker-compose up -d panghostlin-keys
members:
	@-docker-compose stop panghostlin-members
	@-docker-compose build panghostlin-members
	@-docker-compose up -d panghostlin-members
proxy:
	@-docker-compose stop panghostlin-proxy
	@-docker-compose build panghostlin-proxy
	@-docker-compose up -d panghostlin-proxy
pictures:
	@-docker-compose stop panghostlin-pictures
	@-docker-compose build panghostlin-pictures
	@-docker-compose up -d panghostlin-pictures

re:
	docker-compose up -d --build
up:
	docker-compose up --build

purge:
	@-rm -rf src/Pictures
	@-rm -rf src/Members
	@-rm -rf src/Keys
	@-rm -rf src/Proxy
	@-rm -rf src/Webapp
	@-rm -rf .data
	# @-rm -rf .pictures