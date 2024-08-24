package egress

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tuxounet/k-hab/utils"
)

func (h *HttpEgressController) handleOsMirror(w http.ResponseWriter, r *http.Request) {

	h.log.TraceF("\n\rMirror Os Package Request to %s\n\r", r.RequestURI)

	segments := strings.Split(r.RequestURI, "/")
	distro := segments[2]

	mirrors := utils.GetMapValue(h.ctx.GetHabConfig(), "egress.mirrors").(map[string]interface{})

	finalUrl := ""
	for key, mirror := range mirrors {
		if strings.HasPrefix(distro, key) {
			rest := strings.Join(segments[3:], "/")
			finalUrl = fmt.Sprintf("%s%s", mirror, rest)
			break
		}
	}

	if finalUrl == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	resp, err := http.Get(finalUrl)

	if err != nil {
		h.log.DebugF("\nError with egress mirror request to %s: %s\n\r", finalUrl, err.Error())
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for key, value := range resp.Header {
		w.Header().Set(key, value[0])
	}
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)

}
