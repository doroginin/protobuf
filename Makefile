update-swagger:
	go get -u github.com/sergei-svistunov/gostatic2lib

	rm -rf /tmp/swagger-ui
	git clone https://github.com/swagger-api/swagger-ui.git /tmp/swagger-ui

	cd /tmp/swagger-ui; \
		mkdir ./html; \
		cat ./dist/index.html | sed 's/http:\/\/petstore.swagger.io\/v2\///g' > ./html/index.html; \
		cp ./dist/*.js ./html; \
		cp ./dist/*.css ./html; \
		cp ./dist/*.png ./html

	gostatic2lib -out ./protoc-gen-go-http-server/swagger/ui.go -package swaggerui -path /tmp/swagger-ui/html