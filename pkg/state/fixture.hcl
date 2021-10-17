type "nomad-clients" {
  config {
    packages "yum" {}
    services "systemctl" {}
    filesystem "posix" {}
    firewall "iptables" {}
    os "centos7" {}
  }

  system {
    os_option "selinux" {
      value = "disabled"
    }
    user "korchasa" {}
    package "epel-release" {}
    package "yum-utils" {}
    package "unzip" {}
    package "vim-7.29.0-59.el7" {}
    package "noderig" { removed = true }
    package "firewalld" { removed = true }
    file "/home/korchasa/.ssh/authorized_keys" {
      from        = "./korchasa_authorized_keys"
      user        = "korchasa"
      group       = "korchasa"
      permissions = 0600
    }
    file "/etc/yum.repos.d/docker-ce.repo" {
      from        = "https://download.docker.com/linux/centos/docker-ce.repo"
      user        = "root"
      group       = "root"
      permissions = 0644
    }
  }

  application "consul" {
    os_option "env" {
      value = "CONSUL_HTTP_ADDR=http://127.0.0.1:8500"
    }
    user "consul" {
      system = true
    }
    directory "/var/consul/" {
      user        = "consul"
      group       = "consul"
      permissions = 0700
    }
    directory "/etc/consul.d/" {
      user        = "consul"
      group       = "consul"
      permissions = 0700
    }
    file "/usr/local/bin/consul" {
      from        = "https://releases.hashicorp.com/consul/1.9.5/consul_1.9.5_linux_amd64.zip"
      compressed  = true
      user        = "consul"
      group       = "consul"
      permissions = 0600
    }
    file "/etc/consul.d/config.json" {
      from        = "./consul_client_config.json"
      template    = true
      user        = "consul"
      group       = "consul"
      permissions = 0400
    }
    file "/etc/systemd/system/consul.service" {
      from        = "./consul.service"
      user        = "root"
      group       = "root"
      permissions = 0644
    }
    package "unbound" {}
    service "consul" {}
    firewall "consul" {
      ports   = ["8300:8302", "8500:8502", "8600"]
      targets = []
    }
  }

  server "node1-nomad" {
    system {
      os_option "hostname" {
        value = "node1-nomad"
      }
    }
    application "consul" {
      firewall "consul" {
        targets = [
          "144.76.235.93/16",
          "175.148.45.70",
          "4.190.89.53"
        ]
      }
    }
  }
}