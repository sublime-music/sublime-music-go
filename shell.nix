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
    gst
    gst_all_1.gst-plugins-base
    gst_all_1.gst-plugins-good
    gst_all_1.gstreamer
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
