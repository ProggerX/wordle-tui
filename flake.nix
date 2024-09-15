{
    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs";
        flake-parts.url = "github:hercules-ci/flake-parts";
    };
    outputs = { nixpkgs, flake-parts, ... }@inputs: flake-parts.lib.mkFlake { inherit inputs; } {
        systems = nixpkgs.lib.platforms.unix;
        perSystem = { pkgs, ... }: {
			packages.default = import ./default.nix { inherit pkgs; };
        };
    };
}
