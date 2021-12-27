with (import ./nixpkgs.nix);

go.overrideAttrs (old: rec {
  version = "1.18beta1";
  src = fetchurl {
    url = "https://dl.google.com/go/go${version}.src.tar.gz";
    sha256 = "18akwrw4lzl6cj3yy5hrnf4zfycx84xas1s95mdwp6a6n66h5321";
  };
  patches = [];
  doCheck = false;
})
