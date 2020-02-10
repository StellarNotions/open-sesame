# Open Sesame

## Opening the PI to the World

### The Strategy

Because the Raspberry PI is apart of my home network, it is inaccessible by anything that's not connected to the same network.

To allow OpenSesame API to be accessed from anywhere in the world, the following was implemented:

1. Create an AWS EC2 instance with inbound rules set to:

| Type  | Protocol | Port Range | Source |
| ----- | -------- | ---------- | ------ |
| Custom TCP | TCP | 8000 | 0.0.0.0/0 |
| Custom TCP | TCP | 8000 | ::/0 |
| SSH | TCP | 22 | 0.0.0.0/0 |
| SSH | TCP | 22 | ::/0 |

And Outbound rules set to:

| Type  | Protocol | Port Range | Source |
| ----- | -------- | ---------- | ------ |
| All TCP | TCP | 0 - 65535 | 0.0.0.0/0 |
| All TCP | TCP | 0 - 65535 | ::/0 |

2. Create an `autossh` tunnel between the PI and the EC2 instance

### The Execution

Running the following command will allow the Raspberry PI to be accessed from anywhere.

- `autossh -N -R *:8000:127.0.0.1:8000 user@publically-accessible-ip -i "foo.pem"`

However, for Ubuntu servers, the following command will automatically bind port `8000` to the local IP, so `127.0.0.1` instead of the publically accessible IP.

To get around this, you can enable the `GatewayPorts` option in `/etc/ssh/sshd_config`:

```bash
...
#AllowAgentForwarding yes
#AllowTcpForwarding yes
GatewayPorts yes
X11Forwarding yes
#X11DisplayOffset 10
#X11UseLocalhost yes
#PermitTTY yes
...
```
