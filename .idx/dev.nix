{ pkgs, ... }: {
  # 1. Channel
  channel = "stable-24.05";

  # 2. Cài đặt các gói công cụ
  packages = [
    pkgs.go
    pkgs.docker
    pkgs.docker-compose
    pkgs.gnumake
    pkgs.protobuf
    pkgs.evans
  ];

  # 3. KÍCH HOẠT DOCKER (Đặt ở đây, ngoài block idx)
  services.docker.enable = true;

  # 4. Các cấu hình riêng cho IDX
  idx = {
    extensions = [
      "golang.Go"
      "ms-azuretools.vscode-docker"
    ];

    previews = {
      enable = true;
      previews = {
        # web = {
        #   command = ["go" "run" "main.go"];
        #   manager = "web";
        #   env = { PORT = "$PORT"; };
        # };
      };
    };
  };
}