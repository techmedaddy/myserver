# 🔥 YEH KYA BAWAL CHEEZ HAI BHAI 🔥

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/FIELDING-SET_HAI-success?style=for-the-badge)
![Vibes](https://img.shields.io/badge/VIBES-IMMACULATE-purple?style=for-the-badge)

> *"kamal aadmi ho yaar"* - literally everyone who sees this code

## 💀 TL;DR (attention span = 0.2 seconds)

**MAAL CHEEZ HAI** - A chill HTTP server for the homies. No BS, no bloat, just pure **FIELDING SET HAI** energy.

Standard lib only. Clean architecture. Logs that actually make sense. Built different fr fr 💯

## 🎯 Features That Go HARD

- ✨ **Standard Library HTTP Server** - no cap, pure stdlib supremacy
- 🎨 **Router + Handler Separation** - architecture cleaner than your git history
- 🧠 **Service Layer** - business logic stays bussin'
- 📝 **Logger Util** - stdout/stderr pe proper logs, not your console.log spam
- 📁 **Modular Layout** - `cmd/`, `internal/`, `pkg/` - organized like a pro
- 🚀 **Zero Dependencies** - lightweight king behavior

## 🏗️ Project Structure (Architect Mode Activated)

```
myserver/
├── cmd/
│   └── server/
│       └── main.go          // entry point, server ka baap
├── internal/
│   ├── handlers/
│   │   └── handlers.go      // HTTP handlers, requests ko sambhalta hai
│   ├── services/
│   │   └── service.go       // business logic, dimag yahan hai
│   └── router/
│       └── router.go        // routes ka boss
└── pkg/
    └── logger/
        └── logger.go        // logging ka don, stdout/stderr master
```

**MATLAB** - everything in its place, no spaghetti code allowed 🚫🍝

## 🚀 Quick Start (Speedrun Any%)

```bash
# clone kar pehle
git clone <your-repo>
cd myserver

# dependencies? bhai we don't do that here
go mod tidy

# run karke dekh magic
go run cmd/server/main.go

# build karke flex kar
go build -o bin/server cmd/server/main.go
./bin/server
```

**BOOM** 💥 Server running on `localhost:8080` - absolutely legendary

## 🎮 How To Use (Noob-Friendly Guide)

### Basic Server Setup

```go
// cmd/server/main.go
package main

import (
    "log"
    "net/http"
    "myserver/internal/router"
    "myserver/pkg/logger"
)

func main() {
    // logger setup - clean logs unlocked
    logger.Init()
    
    // router setup - routes ka network
    r := router.NewRouter()
    
    logger.Info("Server chal gaya port 8080 pe! 🔥")
    
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal("Server mar gaya bhai: ", err)
    }
}
```

### Handler Example (Copy-Paste Approved ✅)

```go
// internal/handlers/handlers.go
package handlers

import (
    "encoding/json"
    "net/http"
    "myserver/internal/services"
    "myserver/pkg/logger"
)

type Handler struct {
    service *services.Service
}

func NewHandler(s *services.Service) *Handler {
    return &Handler{service: s}
}

func (h *Handler) GetData(w http.ResponseWriter, r *http.Request) {
    logger.Info("Request aaya boss!")
    
    data := h.service.FetchData()
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
    
    logger.Info("Response bhej diya, khatam!")
}
```

### Service Layer (Business Logic Zone 🧠)

```go
// internal/services/service.go
package services

type Service struct {
    // your dependencies here
}

func NewService() *Service {
    return &Service{}
}

func (s *Service) FetchData() map[string]interface{} {
    // actual business logic yahan
    return map[string]interface{}{
        "status": "absolutely legendary",
        "vibe":   "immaculate",
        "energy": "unmatched",
    }
}
```

### Logger (Clean Logs Supremacy 📝)

```go
// pkg/logger/logger.go
package logger

import (
    "log"
    "os"
)

var (
    infoLog  *log.Logger
    errorLog *log.Logger
)

func Init() {
    infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(v ...interface{}) {
    infoLog.Println(v...)
}

func Error(v ...interface{}) {
    errorLog.Println(v...)
}
```

### Router Setup (Traffic Police 🚦)

```go
// internal/router/router.go
package router

import (
    "net/http"
    "myserver/internal/handlers"
    "myserver/internal/services"
)

func NewRouter() *http.ServeMux {
    mux := http.NewServeMux()
    
    // service aur handler setup
    service := services.NewService()
    handler := handlers.NewHandler(service)
    
    // routes - clean aur organized
    mux.HandleFunc("/", handler.GetData)
    mux.HandleFunc("/health", healthCheck)
    
    return mux
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("fielding set hai boss ✅"))
}
```

## 🎪 API Endpoints (Hit These Routes)

| Endpoint | Method | Kya Karta Hai? | Response |
|----------|--------|----------------|----------|
| `/` | GET | Main data fetch karta hai | JSON maal |
| `/health` | GET | Server alive hai check karo | "fielding set hai" |
| `/api/*` | * | Custom routes yahan | Tumhara code |

## 🛠️ Development Tips (Pro Gamer Moves)

### Hot Reload (Developer Comfort Mode)

```bash
# air install karo (optional but based)
go install github.com/cosmtrek/air@latest

# ab har change pe auto reload
air
```

### Testing (Quality Assurance Activated)

```bash
# tests run karo
go test ./...

# coverage dekho
go test -cover ./...

# detailed report - full stats
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🎯 Best Practices (Sigma Grindset)

1. **Separation of Concerns** - handlers ko handlers rehne do, logic service mein
2. **Error Handling** - errors ko properly handle karo, production mein mat phasao
3. **Logging** - sab kuch log karo, debugging easy ho jayega
4. **Context Usage** - request context use karo for timeouts/cancellations
5. **Middleware** - authentication, CORS, logging - middleware se karo

## 🚀 Production Deployment (Boss Mode)

```bash
# optimized build
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server cmd/server/main.go

# Docker se deploy (chad move)
docker build -t myserver:latest .
docker run -p 8080:8080 myserver:latest

# systemd service (linux server vibes)
sudo systemctl enable myserver
sudo systemctl start myserver
```

### Sample Dockerfile

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server cmd/server/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

## 🤝 Contributing (Pull Requests Welcome)

Code clean rakhna, comments likhna, tests likhna. Bas itna yaad rakhna:

1. Fork karo repo
2. Feature branch banao (`git checkout -b feature/insane-feature`)
3. Changes commit karo (`git commit -m 'Added some insane stuff'`)
4. Branch push karo (`git push origin feature/insane-feature`)
5. Pull Request kholo aur chill karo

## 📜 License

MIT - **MAAL HAI** - free for everyone

## 💬 Meme Section (Mandatory)

```
    ∧＿∧
  （｡･ω･｡)つ━☆・*。
  ⊂　 ノ 　　　・゜+.
   しーＪ　　　°。+ *´¨)
              .· ´¸.·*´¨) ¸.·*¨)
              (¸.·´ (¸.·'* ☆ FIELDING SET HAI
```

## 🔥 Final Words

**YEH PROJECT MAAL CHEEZ HAI BHAI** 

No cap, straight fire, production ready. Clone karo, run karo, flex karo 💯

Built with:
- ☕ Chai (obviously)
- 💻 VS Code (IDE supremacy)
- 🎵 Lofi beats (focus mode)
- 😤 Pure dedication
- 🔥 Absolutely legendary energy

---

⭐ **Star the repo if this slaps** ⭐

*kamal aadmi ho yaar for reading this far*

**"Code clean rakhna, bugs dur rakhna, vibes high rakhna"** - Ancient Go Proverb

Made with 💀 by the homies, for the homies

**NO 🧢 DETECTED** ✅