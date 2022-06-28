package handlers

import (
	"context"
	dragonSpider "github.com/kaliadmen/dragon_spider"
	"net/http"
)

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, template string, vars, data any) error {
	return h.App.Render.Page(w, r, template, vars, data)
}

func (h *Handlers) sessionGet(ctx context.Context, key string) any {
	return h.App.Session.Get(ctx, key)
}

func (h *Handlers) sessionPut(ctx context.Context, key string, value any) {
	h.App.Session.Put(ctx, key, value)
}

func (h *Handlers) sessionHas(ctx context.Context, key string) bool {
	return h.App.Session.Exists(ctx, key)
}

func (h *Handlers) sessionRemove(ctx context.Context, key string) {
	h.App.Session.Remove(ctx, key)
}

func (h *Handlers) sessionRenew(ctx context.Context) error {
	err := h.App.Session.RenewToken(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handlers) sessionDestroy(ctx context.Context) error {
	err := h.App.Session.Destroy(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handlers) writeJSON(w http.ResponseWriter, status int, data any) error {
	err := h.App.WriteJSON(w, status, data)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handlers) writeXML(w http.ResponseWriter, status int, data any) error {
	err := h.App.WriteXML(w, status, data)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handlers) fileDownLoad(w http.ResponseWriter, r *http.Request, pathToFile, filename string) error {
	err := h.App.FileDownload(w, r, pathToFile, filename)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handlers) encrypt(text string) (string, error) {
	enc := dragonSpider.Encryption{Key: []byte(h.App.EncryptionKey)}

	encrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func (h *Handlers) decrypt(text string) (string, error) {
	enc := dragonSpider.Encryption{Key: []byte(h.App.EncryptionKey)}

	decrypted, err := enc.Decrypt(text)
	if err != nil {
		return "", err
	}

	return decrypted, nil
}
