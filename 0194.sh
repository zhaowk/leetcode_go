#!/bin/bash

COLS=$(head -1 file.txt | awk '{print NF}')
for ((i=1;i<=COLS;i++)) do
   echo $(awk 'BEGIN{ORS=" "}{print $i}' i=$i file.txt | sed 's/.$//')
done