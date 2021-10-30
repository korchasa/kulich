spec "nomad-clients" {
  config {
    packages "yum" {}
    services "systemctl" {}
    filesystem "posix" {}
    firewall "iptables" {}
    os "centos7" {}
  }

  system {
    os_option "hostnamectl" "hostname" {
      value = "node1-nomad"
    }
    os_option "selinux" "enabled" {
      value = "false"
    }
    user "alice" {}
    package "epel-release" {}
    package "yum-utils" {}
    package "unzip" {}
    package "vim-7.29.0-59.el7" {}
    package "noderig" { removed = true }
    package "firewalld" { removed = true }
    file "/home/korchasa/.ssh/authorized_keys" {
      from        = "./korchasa_authorized_keys"
      user        = "alice"
      group       = "alice"
      permissions = 0600
    }
    file "/etc/yum.repos.d/docker-ce.repo" {
      from        = "./docker-ce.repo"
      user        = "consul"
      group       = "consul"
      permissions = 0600
    }
  }

  application "consul" {
    os_option "env" "CONSUL_HTTP_ADDR" {
      value = "http://127.0.0.1:8500"
    }
    user "consul" {
      system = true
    }
    directory "/var/consul/" {
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
      from          = "./consul_client_config.json"
      template      = true
      template_vars = {
        foo : "bar",
        baz : "42"
      }
      user          = "consul"
      group         = "consul"
      permissions   = 0400
    }
    package "unbound" {}
    service "consul" {

    }
    firewall "consul" {
      ports   = ["8300:8302", "8500:8502", "8600"]
      targets = [
        "444.444.444.444/16",
        "555.555.555.555",
        "666.666.666.66"
      ]
    }
  }
}
