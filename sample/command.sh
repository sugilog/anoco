type=$1

if [ -f sample/command.$type.json ];
then
	bash -c -v "msgpack-cli encode sample/command.$type.json > /tmp/anoco.bin"
	bash -c -v "curl -D /dev/stderr localhost:15525/job -XPOST -d@/tmp/anoco.bin" | tee /tmp/anoco.response.bin
  bash -c -v "msgpack-cli decode /tmp/anoco.response.bin"
else
  echo 'Invalid type given'
  exit 1
fi
