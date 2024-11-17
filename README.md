# k-hab

A single executable, that spawn or restore a virtual infrastructure, definied by plain yaml files.

> [ChangeLog](./CHANGELOG.md)

## Get Release

### Requirements

- snap
- sudo access without password

### Get A Release

```bash
RELEASE=24.11.1
curl -s https://api.github.com/repos/tuxounet/k-hab/releases/tags/${RELEASE} \
| grep "browser_download_url.*" \
| cut -d : -f 2,3 \
| tr -d '"' \
| wget -qi - \
&& chmod +x k-hab-linux-amd64 \
&& mv k-hab-linux-amd64 k-hab \
&& ./k-hab shell
```

## Usage example

> [Examples](./examples/README.md)

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
