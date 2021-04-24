#================================
#== DOCKER ENVIRONMENT
#================================
COMPOSE := @docker-compose

dcb:
	${COMPOSE} build

dcuf:
ifdef f
	${COMPOSE} up -d --${f}
endif

dcu:
	${COMPOSE} up -d --build

dcd:
	${COMPOSE} down

#================================
#== GOLANG ENVIRONMENT
#================================
GO := @go
GIN := @gin

goinstall:
	${GO} get .

godev:
	${GIN} -a 4000 -p 3001 -b bin/gin-bin run server.go

goprod:
	${GO} build -o main .

gotest:
	${GO} test