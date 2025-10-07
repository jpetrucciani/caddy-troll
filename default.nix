{ pkgs ? import
    (fetchTarball {
      name = "jpetrucciani-2025-10-07";
      url = "https://github.com/jpetrucciani/nix/archive/15d79d49616d420eb45e52479c42d57ff8f58537.tar.gz";
      sha256 = "1z373gnlz41zvqjl8hq7ks2nzsss6c1q8mv95vamxzhq6jcsqwfj";
    })
    { }
}:
let
  name = "caddy-troll";
  tools = with pkgs; {
    cli = [
      jfmt
      nixup
      (writeShellScriptBin "test_actions" ''
        export DOCKER_HOST=$(${pkgs.docker-client}/bin/docker context inspect --format '{{.Endpoints.docker.Host}}')
        ${pkgs.act}/bin/act --container-architecture linux/amd64 -r --rm
      '')
    ];
    go = [
      go
      go-tools
      gopls
    ];
    scripts = pkgs.lib.attrsets.attrValues scripts;
  };

  scripts = with pkgs; let
    run-troll = pog {
      name = "run-troll";
      description = "run caddy with the troll plugin in watch mode against the caddyfile in the conf dir";
      script = h: with h; ''
        ${xcaddy}/bin/xcaddy run -- --config ./conf/Caddyfile --watch "$@"
      '';
    };
  in
  {
    inherit run-troll;
    run = pog {
      name = "run";
      description = "run run-troll, restarting when go files are changed";
      script = h: with h; ''
        # wget -nc -q "https://raw.githubusercontent.com/minimaxir/big-list-of-naughty-strings/master/blns.json"
        ${findutils}/bin/find . -iname '*.go' | ${entr}/bin/entr -rz ${run-troll}/bin/run-troll
      '';
    };
  };
  paths = pkgs.lib.flatten [ (builtins.attrValues tools) ];
  env = pkgs.buildEnv {
    inherit name paths; buildInputs = paths;
  };
in
(env.overrideAttrs (_: {
  inherit name;
  NIXUP = "0.0.10";
})) // { inherit scripts; }
