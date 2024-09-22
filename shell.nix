{
  pkgs ? import <nixpkgs> {}
}:
  pkgs.mkShell {
    LD_LIBRARY_PATH = pkgs.lib.makeLibraryPath [
      pkgs.stdenv.cc.cc
    ];

    LOCALE_ARCHIVE = "${pkgs.glibcLocales}/lib/locale/locale-archive";

    buildInputs = [
      pkgs.delve
      pkgs.git
      pkgs.glibcLocales
      pkgs.go_1_22
      pkgs.gofumpt
      pkgs.golangci-lint
      pkgs.gopls
      pkgs.nix
      pkgs.nodePackages.pnpm
      pkgs.nodejs-18_x
    ];

    hardeningDisable = [ "fortify" ];

    shellHook = ''
    '';
  }
