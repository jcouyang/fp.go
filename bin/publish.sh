#/usr/bin/env bash
set -e

VERSION=$(dhall text <<< "./renderVersion.dhall")
echo publishing $VERSION...

gh release create "$VERSION"

GOPROXY=proxy.golang.org go list -m "oyanglul.us/fizpop@$VERSION"
