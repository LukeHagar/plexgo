.PHONY: specs
specs:	
	git clone https://github.com/LukeHagar/plex-api-spec.git api-specs

.PHONY: clean-specs
clean-specs:	
	rm -rf ./api-specs

build:	
	java -jar openapi-generator-cli.jar generate -i api-specs/pms/pms-spec.yaml -g go -o pms --global-property skipFormModel=false --config sdk-resources/pms-config.yaml -p enumClassPrefix=true --git-repo-id plexgo --git-user-id lukehagar
	node sdk-resources/postscript.js ./pms
	java -jar openapi-generator-cli.jar generate -i api-specs/plextv/plextv-spec.yaml -g go -o plextv --global-property skipFormModel=false --config sdk-resources/plextv-config.yaml -p enumClassPrefix=true --git-repo-id plexgo --git-user-id lukehagar
	node sdk-resources/postscript.js ./plextv
	