{ jacobi ? import
    (
      fetchTarball {
        name = "jpetrucciani-2022-10-04";
        url = "https://github.com/jpetrucciani/nix/archive/05cf540fbb9784c4ebec45190073efaf82e45fd7.tar.gz";
        sha256 = "19qp5rd1360j9vhkk86xjwcy6j45qik4nxgy61cl2m4591fkjwl7";
      }
    )
    { }
}:
let
  name = "caddy-troll";
  tools = with jacobi; {
    cli = [
      nixpkgs-fmt
    ];
    go = [
      go_1_19
      go-tools
      gopls
    ];
    scripts = [
      xcaddy
      (jacobi.pog {
        name = "run-troll";
        script = h: with h; ''
          ${xcaddy}/bin/xcaddy run --config ./conf/Caddyfile --watch
        '';
      })
    ];
  };

  env = jacobi.enviro {
    inherit name tools;
  };
in
env
