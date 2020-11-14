make build
$(GCC) $(GCMD) $(GPATH)

build :
	make build_db

build_db :
	rm pkg/db/db_structs.go
	go run pkg/main.go -json=./pkg/db/config.json
	mv db_structs.go pkg/db/

install :
	make install_routes
	make install_db
	make install_auth
	make install_encryption
	
install_encryption :
	go get -u golang.org/x/crypto/bcrypt
install_routes :
	go get -u github.com/gorilla/mux
install_db :
	go get -u github.com/go-xorm/xorm
install_auth :
	go get github.com/dgrijalva/jwt-go

# install_go

create_keys:
	ssh-keygen -m PEM -t rsa -b 4096 -f keys/app.rsa
	openssl rsa -in keys/app.rsa -pubout -outform PEM -out keys/app.rsa.pub