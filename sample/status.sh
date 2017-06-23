bash -c -v "curl -D /dev/stderr \"localhost:15525/job/status?uuid=$1\"" | tee /tmp/anoco.status.bin
bash -c -v "msgpack-cli decode /tmp/anoco.status.bin"
