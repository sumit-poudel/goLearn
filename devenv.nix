{
  pkgs,
  lib,
  config,
  inputs,
  ...
}:
{
  # https://devenv.sh/languages/
  languages.go.enable = true;
  packages = with pkgs; [
    gopls
    golangci-lint
  ];
}
