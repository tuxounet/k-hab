# k-hab 24.9.0

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


# Based on :

- Linux
- snap
- Incus ecosystem :
  - IncusD Service
  - incus client
  - Distrobuilder
- Alpine Linux

# Built with :

- Go Lang 1.22+
