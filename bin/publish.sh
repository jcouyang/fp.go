#/usr/bin/env bash
set -e

VERSION=$(dhall text <<< "./renderVersion.dhall")
echo publishing $VERSION...

LAST_RELEASE=$(gh release view --json tagName -q .tagName)

gh release create "$VERSION" --notes "$(git log "${LAST_RELEASE}"..HEAD --pretty="format:* %s")"

GOPROXY=proxy.golang.org go list -m "github.com/jcouyang/fizpop@$VERSION"
