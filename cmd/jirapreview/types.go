package jirapreview

import "github.com/sirupsen/logrus"

type JiraPreview struct {
	config Config
	logger *logrus.Logger
}

type LinkPreview struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Response struct {
	LinkPreviews map[string]LinkPreview `json:"link_previews"`
}

type Link struct {
	Url    string `json:"url"`
	Domain string `json:"domain"`
}

type Request struct {
	Type       string `json:"type"`
	Event      string `json:"event"`
	Chat_id    int    `json:"chat_id"`
	Message_id int    `json:"message_id"`
	Links      []Link `json:"links"`
}

type Fields struct {
	Priority  Priority  `json:"priority"`
	Project   Project   `json:"project"`
	Summary   string    `json:"summary"`
	Assignee  Assignee  `json:"assignee"`
	Status    Status    `json:"status"`
	IssueType IssueType `json:"issuetype"`
}

type Priority struct {
	Name string `json:"name"`
}

type Assignee struct {
	DisplayName string `json:"displayName"`
}

type Status struct {
	Name string `json:"name"`
}

type IssueType struct {
	Name string `json:"name"`
}

type Project struct {
	Name string `json:"name"`
}

type Issue struct {
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

// RESPONSE
// {
//     "link_previews": {
//       "https://example.com/4176": {
//         "title": "Статья: Отправка файлов",
//         "description": "Пример отправки файлов на удаленный сервер",
//         "image": {
//           "key": "attaches/files/93746/e354fd79-9jh6-f2hd-fj83-709dae24c763/${filename}",
//           "name": "files-to-server.jpg",
//           "size": "695604"
//         }
//       },
//       "https://tasks.example.com/l/38765": {
//         "title": "Задача: Верстка сайта",
//         "description": "Необходимо сверстать одностраничный сайт по макету",
//         "image_url": "https://website.com/img/landing.png"
//       }
//     }
//   }

// REQUEST
// {
//     "type": "message", //тип объекта
//     "event": "link_shared", //тип события
//     "chat_id": 49218, //идентификатор чата, в котором находится сообщение
//     "message_id": 84721, //идентификатор сообщения, к которому относятся ссылки
//     "links": [ //массив объектов найденых ссылок
//       {
//         "url": "https://plus.google.com", //полная ссылка в сообщении
//         "domain": "plus.google.com" //домен, который был найден в сообщении
//       },
//       {
//         "url": "https://meet.google.com/s/123",
//         "domain": "meet.google.com"
//       },
//       {
//         "url": "https://google.com",
//         "domain": "google.com"
//       }
//     ]
//   }
