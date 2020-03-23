#!/bin/bash
echo "CATHOLiC BiBLE TOOLS"
cd dist
function package_app {
	echo "Building package for '$2-$3' ..."
    cd $2-$3
    zip -r ../cbt-$1-$2-$3.zip ./* >/dev/null 2>&1
	cd ..
}
[[ -z "$1" ]] && {
	echo "Error: A version number is required to build packages (ex. 'package.sh 1.00').";
	exit 1;
}
package_app $1 "linux" "386"
package_app $1 "linux" "amd64"
package_app $1 "macos" "386"
package_app $1 "macos" "amd64"
package_app $1 "windows" "386"
package_app $1 "windows" "amd64"
