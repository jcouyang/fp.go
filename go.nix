with (import ./nixpkgs.nix);

go.overrideAttrs (old: rec {
  version = "1.18rc1";
  src = fetchurl {
     url = "https://dl.google.com/go/go${version}.src.tar.gz";
    sha256 = "1f0bk11qra05ywv7zw82gadjigipvrjkdcr1i1gsi3q0adk7mv2w";
  };
  patches = [];
  doCheck = false;
})
