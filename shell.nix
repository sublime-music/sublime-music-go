{ forCI ? false }:
let
  pkgs = import <nixpkgs> { };
in
with pkgs;
mkShell {
  buildInputs = [
    # Core GTK
    gdk-pixbuf
    glib
    go
    gobject-introspection
    graphene
    gtk3
    gtk4
    libadwaita

    # Gstreamer for audio
    gst_all_1.gstreamer
    gst_all_1.gst-plugins-base
    (gst_all_1.gst-plugins-good.override { gtkSupport = true; })
    gst_all_1.gst-plugins-bad
    gst_all_1.gst-plugins-ugly
    gst_all_1.gst-libav

    pkgconfig
    vulkan-headers
  ] ++ lib.lists.optional (!forCI) [
    gotools
    gopls
    vgo2nix
    yq-go
  ];
}
