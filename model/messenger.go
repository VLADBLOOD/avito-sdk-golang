package model

// ChatsListV2Request - запрос списка чатов (v2 API)
type ChatsListV2Request struct {
	ItemIDs    []int64  `json:"item_ids,omitempty"`
	UnreadOnly bool     `json:"unread_only,omitempty"`
	ChatTypes  []string `json:"chat_types,omitempty"`
	Limit      int      `json:"limit,omitempty"`
	Offset     int      `json:"offset,omitempty"`
}

// ChatsListV2Response - ответ API для списка чатов (v2 API)
type ChatsListV2Response struct {
	Chats []ChatInfo `json:"chats"`
}

// ChatInfo - информация о чате
type ChatInfo struct {
	ID          string      `json:"id"`
	Created     int64       `json:"created"`
	Updated     int64       `json:"updated"`
	Context     ChatContext `json:"context"`
	LastMessage Message     `json:"last_message,omitempty"`
	Users       []ChatUser  `json:"users"`
}

// ChatContext - контекст чата
type ChatContext struct {
	Type  string       `json:"type"`
	Value ContextValue `json:"value"`
}

// ContextValue - значение контекста чата
type ContextValue struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	PriceString string     `json:"price_string,omitempty"`
	URL         string     `json:"url"`
	UserID      int64      `json:"user_id"`
	Images      ItemImages `json:"images,omitempty"`
	StatusID    int        `json:"status_id,omitempty"`
}

// ItemImages - изображения объявления
type ItemImages struct {
	Count int               `json:"count"`
	Main  map[string]string `json:"main,omitempty"`
}

// Message - сообщение в чате
type Message struct {
	ID        string         `json:"id"`
	AuthorID  int64          `json:"author_id"`
	Created   int64          `json:"created"`
	Direction string         `json:"direction"`
	Type      string         `json:"type"`
	Content   MessageContent `json:"content"`
	IsRead    bool           `json:"is_read,omitempty"`
	Read      int64          `json:"read,omitempty"`
	Quote     QuotedMessage  `json:"quote,omitempty"`
}

// MessageContent - содержимое сообщения
type MessageContent struct {
	Text     string          `json:"text,omitempty"`
	Call     CallContent     `json:"call,omitempty"`
	Image    ImageContent    `json:"image,omitempty"`
	Item     ItemContent     `json:"item,omitempty"`
	Link     LinkContent     `json:"link,omitempty"`
	Location LocationContent `json:"location,omitempty"`
	Voice    VoiceContent    `json:"voice,omitempty"`
	FlowID   string          `json:"flow_id,omitempty"`
}

// CallContent - содержимое звонка в сообщении
type CallContent struct {
	Status       string `json:"status"`
	TargetUserID int64  `json:"target_user_id"`
}

// ImageContent - изображение в сообщении
type ImageContent struct {
	Sizes map[string]string `json:"sizes,omitempty"`
}

// ItemContent - объявление в сообщении
type ItemContent struct {
	Title       string `json:"title"`
	PriceString string `json:"price_string,omitempty"`
	ItemURL     string `json:"item_url"`
	ImageURL    string `json:"image_url,omitempty"`
}

// LinkContent - ссылка в сообщении
type LinkContent struct {
	Text    string      `json:"text"`
	URL     string      `json:"url"`
	Preview LinkPreview `json:"preview,omitempty"`
}

// LinkPreview - превью ссылки
type LinkPreview struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Domain      string            `json:"domain"`
	URL         string            `json:"url"`
	Images      map[string]string `json:"images,omitempty"`
}

// LocationContent - локация в сообщении
type LocationContent struct {
	Kind  string  `json:"kind"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
	Text  string  `json:"text"`
	Title string  `json:"title"`
}

// VoiceContent - голосовое сообщение
type VoiceContent struct {
	VoiceID string `json:"voice_id"`
}

// QuotedMessage - цитируемое сообщение
type QuotedMessage struct {
	ID       string         `json:"id"`
	AuthorID int64          `json:"author_id"`
	Created  int64          `json:"created"`
	Type     string         `json:"type"`
	Content  MessageContent `json:"content"`
}

// ChatUser - пользователь в чате
type ChatUser struct {
	ID                int64       `json:"id"`
	Name              string      `json:"name"`
	PublicUserProfile UserProfile `json:"public_user_profile,omitempty"`
}

// UserProfile - публичный профиль пользователя
type UserProfile struct {
	UserID int64      `json:"user_id"`
	ItemID int64      `json:"item_id"`
	URL    string     `json:"url"`
	Avatar UserAvatar `json:"avatar,omitempty"`
}

// UserAvatar - аватар пользователя
type UserAvatar struct {
	Default string            `json:"default"`
	Images  map[string]string `json:"images,omitempty"`
}

// ChatMessagesListV3Request - запрос списка сообщений чата (v3 API)
type ChatMessagesListV3Request struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

// ChatMessagesListV3Response - ответ API для списка сообщений чата (v3 API)
type ChatMessagesListV3Response struct {
	Messages []Message `json:"messages"`
}
