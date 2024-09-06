# k-hab

A single executable, that spawn or restore a virtual infrastructure, definied by plain yaml files.

## Get Release

** LINUX amd64 **

```bash
curl -s https://api.github.com/repos/tuxounet/k-hab/releases/tags/24.8.7 \
| grep "browser_download_url.*" \
| cut -d : -f 2,3 \
| tr -d '"' \
| wget -qi - \
&& chmod +x k-hab-linux-amd64 \
&& mv k-hab-linux-amd64 k-hab \
&& ./k-hab shell
```

explore the ".

# Based on :

- Linux
- snap
- LXD ecosystem :
  - LXD Deamon
  - LXC (from LXD)
  - Distrobuilder
- Alpine Linux

# Built with :

- Go Lang 1.22+
