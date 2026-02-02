{
  description = "Development environment for neet-code-150";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";
  };

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" ];
      perSystem = { config, self', inputs', pkgs, system, ... }: {
        devShells.default = pkgs.mkShell {
          name = "default-dev-shell";
          meta.description = "Default development shell";
          packages = with pkgs; [
            go
            gopls
            delve
            golangci-lint
            gotools
         ];
        };
      };
    };
}
