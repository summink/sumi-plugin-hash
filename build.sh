platforms=$(jq -r '.platforms[] | "\(.os)-\(.arch)"' manifest.json)

for platform in $platforms
do
  IFS="/" read -r GOOS GOARCH <<< "$platform"
  GOOS=$GOOS GOARCH=$GOARCH go build -o "$output_name" "$plugin"
done
