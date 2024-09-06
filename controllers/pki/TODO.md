[ ] 1. create a CA , local, based on hab config and this lib
go get github.com/kairoaraujo/goca
https://github.com/kairoaraujo/goca?tab=readme-ov-file#goca-http-rest-api
[ ] 2. Create server ssl keypairs, signed by ca

[ ] 3. create client certificate pairs for mtls server connection establish, else, 404 in ingress

[ ] 4. Certificate can be downloaded , but only from specific ips, maybe as bundle (with ca and client cert)
