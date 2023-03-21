FROM ghcr.io/vmware-labs/httpd-mod-wasm:latest
COPY conf/httpd.conf /usr/local/apache2/conf/httpd.conf
COPY wasm_modules/uuid-cgi.wasm /usr/local/apache2/wasm_modules/
