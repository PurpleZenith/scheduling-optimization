for n in {1..200};
do
   echo "reversing" $n
   tac iter3_$n > reverse_iter3_$n
done