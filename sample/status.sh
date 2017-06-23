bash -c -v "curl -D /dev/stderr \"localhost:15525/job/status?uuid=$1\" | msgpack-cli decode"
