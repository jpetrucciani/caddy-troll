{ jacobi ? import
    (fetchTarball {
      name = "jpetrucciani-2022-10-29";
      url = "https://nix.cobi.dev/x/226c57d8dceeb0556b5405ccc674f24e2c97307b";
      sha256 = "068i7zpvgk9lydbhksyqr58vildjwmp01rhvd2r9fh61sbplpbmj";
    })
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
          # wget -nc -q "https://raw.githubusercontent.com/minimaxir/big-list-of-naughty-strings/master/blns.json"
          ${findutils}/bin/find . -iname '*.go' | ${entr}/bin/entr -rz ${run-troll}/bin/run-troll
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
        (writeShellScriptBin "test_actions" ''
          export DOCKER_HOST=$(${jacobi.docker-client}/bin/docker context inspect --format '{{.Endpoints.docker.Host}}')
          ${jacobi.act}/bin/act --container-architecture linux/amd64 -r --rm
        '')
      ];
    };

  env = jacobi.enviro {
    inherit name tools;
  };
in
env
