type=$1

if [ -f sample/command.$type.json ];
then
	bash -c -v "msgpack-cli encode sample/command.$type.json | curl -D /dev/stderr localhost:15525/job -XPOST -d@/tmp/anoco.bin | msgpack-cli decode"
else
  echo 'Invalid type given'
  exit 1
fi
