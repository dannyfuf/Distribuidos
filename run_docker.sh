# check if $1 is 17 then print "a", if is 18 then print "b", if is 19 then print "c", if is 20 then print "d"
echo "Running machine: $1"
if [ $1 -eq 17 ]
then
    docker exec -it maquina17 /bin/bash
elif [ $1 -eq 18 ]
then
    docker exec -it maquina18 /bin/bash
elif [ $1 -eq 19 ]
then
    docker exec -it maquina19 /bin/bash
elif [ $1 -eq 20 ]
then
    docker exec -it maquina20 /bin/bash
else
    echo "Machine not found"
fi