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
  tools = with jacobi;
    let
      run-troll = pog {
        name = "run-troll";
        description = "run caddy with the troll plugin in watch mode against the caddyfile in the conf dir";
        script = h: with h; ''
          ${xcaddy}/bin/xcaddy run --config ./conf/Caddyfile --watch "$@"
        '';
      };
      run = pog {
        name = "run";
        description = "run run-troll, restarting when go files are changed";
        script = h: with h; ''
          wget -nc -q "https://raw.githubusercontent.com/minimaxir/big-list-of-naughty-strings/master/blns.json"
          ${watchexec}/bin/watchexec -r -e go -- ${run-troll}/bin/run-troll
        '';
      };
    in
    {
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
        run-troll
        run
      ];
    };

  env = jacobi.enviro {
    inherit name tools;
  };
in
env
