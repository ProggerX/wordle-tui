{ pkgs ? import <nixpkgs> {} }:
pkgs.buildGoModule {
	vendorHash = "sha256-R+NckdjFLR7GGPhprlOAB0zLLw71qgTdra816Z6qGlA=";
	src = ./.;
	pname = "wordle-tui";
	version = "1.0.0";
}
