#!/usr/bin/env bash
version=$1

if [[ -z "$version" ]]; then
  echo "usage: $0 <version>"
  exit 1
fi

platforms=("windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64" "linux/amd64")
package_name="drive-encrypt"

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name'-'$GOOS'-'$GOARCH'-'$version

	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o builds/$output_name $package

    if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi

done
