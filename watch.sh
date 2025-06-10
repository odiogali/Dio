while true; do 
  go build -o _build/Dio && pkill -f "_build/Dio"
  inotifywait -e attrib $(find . -name '*.go') || exit
done
