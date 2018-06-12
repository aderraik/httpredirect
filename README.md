# httpredirect - Simple HTTP to HTTPS Redirector

This is a simple GO lang program to redirect the _HTTP_ connections from port _80_ to _HTTPS_ into port _443_.

```bash
# Control whether service loads on boot
systemctl enable httpredirect

# Manual start and stop
systemctl start httpredirect
systemctl stop httpredirect

# Restarting/reloading
systemctl daemon-reload # Run if .service file has changed
systemctl restart

```
