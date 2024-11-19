ROOT_FOLDER := $(shell  git rev-parse --show-toplevel )
RUN_ARGS := --loglevel=TRACE --setup=setup.hab.yaml
CMD := go run ${ROOT_FOLDER}/main.go 

install:
	$(CMD) install ${RUN_ARGS}

uninstall:
	$(CMD) uninstall ${RUN_ARGS}


provision: 
	$(CMD) provision  ${RUN_ARGS}
	
unprovision:
	$(CMD) unprovision ${RUN_ARGS}
up:
	$(CMD) up ${RUN_ARGS}
deploy:
	$(CMD) deploy ${RUN_ARGS}
shell:
	$(CMD) shell ${RUN_ARGS}
run:
	$(CMD) run ${RUN_ARGS}
undeploy:
	$(CMD) undeploy ${RUN_ARGS}
down:
	$(CMD) down ${RUN_ARGS}
rm:
	$(CMD) rm ${RUN_ARGS}
nuke:
	$(CMD) nuke ${RUN_ARGS}

