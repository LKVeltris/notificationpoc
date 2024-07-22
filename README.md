# notificationpoc

## Install Fluentd:
1. curl -L https://td-toolbelt.herokuapp.com/sh/install-ubuntu-bionic-td-agent3.sh | sh

2. `/etc/fluent/fluent.conf`:

```
<source>
  @type tail
  path /home/vagrant/Github/notificationpoc/app.log
  pos_file /var/log/td-agent/app.log.pos
  tag log.app
  <parse>
    @type none
  </parse>
</source>

<match log.app>
  @type nats
  nats_servers nats://localhost:4222
  subject logs
</match>

```

## Install NATS.io:
sudo docker pull nats
sudo docker run -d --name nats-server -p 4222:4222 -p 8222:8222 nats

## Install fluentd plugin
sudo td-agent-gem install fluent-plugin-nats