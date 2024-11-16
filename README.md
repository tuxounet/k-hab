# k-hab 24.9.0

A single executable, that spawn or restore a virtual infrastructure, definied by plain yaml files.

## Get Release

### Requirements

- snap
- sudo access without password

### Get Latest Release

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

### Tested on

- Ubuntu 24.04.1 amd64

# Based on:

- Linux
- snap
- Distrobuilder
  - Alpine Linux image
  - Ubuntu Linux Image
- LXD / LXC (Incus)

# Built with :

- Go Lang 1.22+
