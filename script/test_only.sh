#bin/bash

# script folder
script_folder=$(dirname $(readlink -f "$0"))

#script folder
projFolder=$(dirname $script_folder)

echo "projFolder: $projFolder"

# if the name of the folder is ConfBackend
#echo $(basename $projFolder)
