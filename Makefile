.PHONY: all swagger modified update update-blueprint update-catalog update-deployment update-iaas clean
SWAGGER_VERSION=0.20.1

all:
	go build -o sdk-test

swagger: check-swagger
	rm -rf pkg/client pkg/models
	./hack/fix_iaas_swagger
	swagger mixin -c=1 swagger/vra-iaas-fixed.json swagger/vra-blueprint.json swagger/vra-catalog.json swagger/vra-deployment.json | python3 -mjson.tool > swagger/vra-combined.json
	./hack/fix_vra_swagger --omit-security
	swagger generate client -f swagger/vra-combined.json -t pkg
	./hack/fixup.sh

check-swagger:
	@swagger version | grep ${SWAGGER_VERSION} > /dev/null || { echo "Wrong version of swagger command. Install go-swagger ${SWAGGER_VERSION}"; exit 1; }

modified:
	git ls-files --modified | xargs git add

update: update-blueprint update-catalog update-deployment update-iaas

update-blueprint:
	curl 'https://api.mgmt.cloud.vmware.com/blueprint/api/swagger/swagger-api-docs?group=2019-09-12' | python3 -mjson.tool > swagger/vra-blueprint.json

update-catalog:
	curl 'https://api.mgmt.cloud.vmware.com/catalog/api/swagger/swagger/v2/api-docs?group=2019-01-15' | python3 -mjson.tool > swagger/vra-catalog.json

update-deployment:
	curl 'https://api.mgmt.cloud.vmware.com/deployment/api/swagger/swagger/v2/api-docs?group=2019-01-15' | python3 -mjson.tool > swagger/vra-deployment.json

update-iaas:
	curl 'https://api.mgmt.cloud.vmware.com/iaas/api/swagger/swagger/v2/api-docs?group=iaas' | python3 -mjson.tool > swagger/vra-iaas.json

clean:
	rm swagger/vra-combined.json
