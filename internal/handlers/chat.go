package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/27vhd/raven-chat/internal/models"
	"github.com/27vhd/raven-chat/internal/repository"
)

type ChatHandler struct {
	Repo            repository.ChatRepository
	indexTemplate   *template.Template
	messageTemplate *template.Template
}

func NewChatHandler(repo repository.ChatRepository) *ChatHandler {
	indexTmpl := template.Must(template.ParseFiles("templates/base.html", "templates/components/message.html"))
	messageTmpl := template.Must(template.ParseFiles("templates/components/message.html"))

	return &ChatHandler{
		Repo:            repo,
		indexTemplate:   indexTmpl,
		messageTemplate: messageTmpl,
	}
}

func (h *ChatHandler) Index(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Repo.GetAllMessages()
	if err != nil {
		log.Printf("Error fetching messages: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := h.indexTemplate.Execute(w, messages); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *ChatHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	content := strings.TrimSpace(r.PostFormValue("content"))
	username := r.PostFormValue("username")

	if content == "" {
		http.Error(w, "Message cannot be empty", http.StatusBadRequest)
		return
	}

	msg := models.Message{
		Username:  username,
		Content:   content,
		Timestamp: time.Now(),
	}

	if err := h.Repo.SaveMessage(msg); err != nil {
		log.Printf("Error saving message: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := h.messageTemplate.Execute(w, msg); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
