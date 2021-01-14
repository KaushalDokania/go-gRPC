rm *.pem
rm *.srl
# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=IN/ST=New-Delhi/L=Delhi/O=Crypto/OU=CA/CN=*.kaushaldokania.com/emailAddress=ca-kaushaldokania@gmail.com"

# echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2.1. Generate server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=IN/CN=server"

# 2.2. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

# 3.1. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=IN/CN=client"

# 3.2. Use CA's private key to sign web client's CSR and get back the signed certificate
openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf

# echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text