{ jacobi ? import
    (
      fetchTarball {
        name = "jpetrucciani-2022-10-03";
        url = "https://github.com/jpetrucciani/nix/archive/e11675e2d2a1484780beacd6c910e0ee5118133a.tar.gz";
        sha256 = "15fdp5jjc776d36yx5h8g6vcfq89yx3zjp3x8pccd4hpb669g60s";
      }
    )
    { }
}:
let
  inherit (jacobi.hax) ifIsLinux ifIsDarwin;

  name = "caddy-troll";
  tools = with jacobi; {
    cli = [
      nixpkgs-fmt
    ];
    go = [
      go_1_19
      go-tools
    ];
  };

  env = jacobi.enviro {
    inherit name tools;
  };
in
env
