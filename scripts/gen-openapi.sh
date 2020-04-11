for directory in $(ls cmd/)
do
	pushd cmd/$directory
	$1 init
	popd > /dev/null
done