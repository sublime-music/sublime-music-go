{ forCI ? false }:
let
  pkgs = import <nixpkgs> { };
in
with pkgs;
mkShell {
  buildInputs = [
    gdk-pixbuf
    glib
    go
    gobject-introspection
    graphene
    gtk3
    gtk4
    pkgconfig
    vulkan-headers
  ] ++ lib.lists.optional (!forCI) [
    gotools
    gopls
    vgo2nix
    yq-go
  ];
}
