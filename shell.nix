{ forCI ? false }:
with import <nixpkgs> { };
let
  projectUris = {
    gotk4 = "git@github.com:diamondburned/gotk4.git";
    gtk4-adwaita = "git@github.com:diamondburned/gotk4-adwaita.git";
    gtk = "https://gitlab.gnome.org/GNOME/libadwaita.git";
    libadwaita = "https://gitlab.gnome.org/GNOME/libadwaita.git";
  };

  cloneCmd = rootDir: key: uri: ''
    if [[ -d ${rootDir}/${key} ]]; then
      echo "${rootDir}/${key} already exists. Will not create."
    else
      mkdir -p ${rootDir}
      ${git}/bin/git clone --recurse-submodules -j8 ${uri} ${rootDir}/${key}
    fi
  '';
  recurseProjectUris = rootDir: lib.mapAttrsToList (
    name: value:
      if builtins.isAttrs value
      then recurseProjectUris "${rootDir}/${name}" value
      else cloneCmd rootDir name value
  );

  PROJECT_ROOT = builtins.getEnv "PWD";

  initGitPkg = writeShellScriptBin "initgit" ''
    echo
    echo Cloning necessary repos
    echo
    ${lib.concatStringsSep "\n" (lib.flatten (recurseProjectUris PROJECT_ROOT projectUris))}
    echo
  '';
in
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
    libxml2

    # Gstreamer for audio
    gst_all_1.gstreamer
    gst_all_1.gst-plugins-base
    (gst_all_1.gst-plugins-good.override { gtkSupport = true; })
    gst_all_1.gst-plugins-bad
    gst_all_1.gst-plugins-ugly
    gst_all_1.gst-libav

    # Build utilities
    gnumake
    pkgconfig
    vulkan-headers
  ] ++ lib.lists.optional (!forCI) [
    gi-docgen
    gopls
    gotools
    initGitPkg
    vgo2nix
    yq-go
  ];

  # Workaround for the lack of wrapGAppsHook:
  # https://nixos.wiki/wiki/Development_environment_with_nix-shell
  shellHook = with pkgs.gnome; with pkgs; ''
    export GDK_PIXBUF_MODULE_FILE='${librsvg}/lib/gdk-pixbuf-2.0/2.10.0/loaders.cache'
    export XDG_DATA_DIRS=$XDG_DATA_DIRS:${hicolor-icon-theme}/share:${adwaita-icon-theme}/share
    export XDG_DATA_DIRS=$XDG_DATA_DIRS:$GSETTINGS_SCHEMAS_PATH
  '';
}
