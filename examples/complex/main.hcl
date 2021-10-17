server "nomad-client" {
  config {
    packages "yum" {}
    services "systemctl" {}
    firewall "iptables" {}
    os "centos7" {
      option {
        type  = "hostname"
        value = "foobar"
      }
      option {
        type  = "selinux"
        value = "disabled"
      }
      option {
        type  = "timezone"
        value = "UTC"
      }
      option {
        type  = "env"
        value = "CONSUL_HTTP_ADDR=http://127.0.0.1:8500"
      }
      option {
        type  = "sysctl"
        value = "net.core.rmem_max=2500000"
      }
    }
  }

  prepare {
    package "epel-release" {}
    package "yum-utils" {}
    package "unzip" {}
    package "libselinux-python" {}
    package "vim-7.29.0-59.el7" {}
    package "tmux" {}
    package "htop" {}
    package "git" {}
    package "zsh" {}
    package "curl" {}
    package "mc" {}
    package "net-tools" {}
    package "iptables" {}
    package "noderig" { removed = true }
    package "firewalld" { removed = true }
    user "korchasa" {}
    file "/home/korchasa/.ssh/authorized_keys" {
      from        = "korchasa_authorized_keys"
      owner       = "korchasa"
      group       = "korchasa"
      permissions = 0600
    }
    file "/etc/yum.repos.d/docker-ce.repo" {
      from        = "https://download.docker.com/linux/centos/docker-ce.repo"
      owner       = "root"
      group       = "root"
      permissions = 0644
    }
  }

  application "yum-cron" {
    package "yum-cron" {}
    file "/etc/yum/yum-cron.conf" {
      from        = "yum-cron.conf"
      owner       = "root"
      group       = "root"
      permissions = 0644
    }
    service "yum-cron" {}
  }

  application "unbound" {
    file "/etc/unbound/unbound.conf" {
      from = "unbound.conf"
    }
    file "/etc/resolv.conf" {
      content = "nameserver 127.0.0.1"
    }
    package "unbound" {}
    service "unbound" {}
  }

  application "node_exporter" {
    user "node_exporter" {
      shell = "/bin/nologin"
      home  = ""
    }
    directory "/etc/node_exporter" {
      touch       = true
      owner       = "root"
      group       = "root"
      permissions = 0644
    }
    file "/usr/local/bin/node_exporter" {
      from_url    = "https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz"
      archived    = true
      owner       = "root"
      group       = "root"
      permissions = 0644
    }
    file "/etc/systemd/system/node_exporter.service" {
      from        = "node_exporter.service"
      owner       = "root"
      group       = "root"
      permissions = 0644
    }
    service "node_exporter" {}
    firewall "node_exporter" {
      port = 9100
      ips  = [
        "144.76.235.93/16",
        "51.178.78.52",
      ]
    }
  }

  application "docker" {
    package "yum-utils" {}
    package "device-mapper-persistent-data" {}
    package "lvm2" {}
    package "libselinux-python" {}
    package "docker-ce" {}
    package "docker-ce-cli" {}
    package "containerd.io" {}
    package "python-pip" {}
    file "/etc/docker/daemon.json" {
      from_path   = "docker.json"
      owner       = "root"
      group       = "root"
      permissions = 0644
    }
    service "docker" {}
  }

  application "consul" {
    user "consul" {
      shell = "/bin/nologin"
      home  = "/var/consul"
    }
    directory "/var/consul/" {
      owner = "consul"
      group = "consul"
    }
    directory "/etc/consul.d/" {
      owner = "consul"
      group = "consul"
    }
    file "/usr/local/bin/consul" {
      from      = "https://releases.hashicorp.com/consul/1.9.5/consul_1.9.5_linux_amd64.zip"
      unarchive = true
    }
    file "/etc/consul.d/config.json" {
      from     = "consul_client_config.json"
      template = true
      owner    = "consul"
      group    = "consul"
      mode     = 0400
    }
    file "/etc/systemd/system/consul.service" {
      from_path   = "consul.service"
      owner       = "root"
      group       = "root"
      permissions = 0644
    }
    service "consul" {}
  }
}
