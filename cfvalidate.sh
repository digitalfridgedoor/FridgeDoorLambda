dir=$(pwd)/$1

echo validating $dir

aws cloudformation validate-template --template-body file://$dir