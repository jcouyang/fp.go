with (import ./nixpkgs.nix);
mkShell {
  shellHook = ''
    export PATH=$HOME/go/bin/:$PATH
  '';
  buildInputs = [
    dhall
    gh
    go
  ];
}
