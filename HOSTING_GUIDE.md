# 🚀 Guia Completo de Hospedagem - Discord Boost Bot

## 📋 Opções de Hospedagem

### 1. **Docker + VPS (Recomendado)** ⭐
### 2. **Heroku** (Gratuito, mas limitado)
### 3. **Railway** (Moderno e fácil)
### 4. **Replit** (Online IDE)
### 5. **AWS/Google Cloud** (Produção)

---

## ✅ Opção 1: VPS + Docker (Recomendado para Produção)

### Provedores:
- **DigitalOcean** ($5-12/mês)
- **Linode** ($5-10/mês)
- **Hetzner** (€3-20/mês)
- **AWS EC2** (t2.micro grátis por 1 ano)

### Passo a Passo:

#### 1. Criar VPS
```bash
# Escolha um SO: Ubuntu 22.04 LTS
# Acesse seu servidor via SSH
ssh root@seu_ip_vps
```

#### 2. Instalar Docker
```bash
# Atualizar sistema
apt update && apt upgrade -y

# Instalar Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

# Instalar Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Verificar instalação
docker --version
docker-compose --version
```

#### 3. Clonar Repositório
```bash
cd /opt
git clone https://github.com/xre2033-hub/discord-boost-bot-automatico.git
cd discord-boost-bot-automatico
```

#### 4. Configurar Variáveis de Ambiente
```bash
# Copiar template
cp .env.example .env

# Editar com seu editor favorito
nano .env
# ou
vi .env
```

Preencha com:
```env
DISCORD_TOKEN=seu_token_aqui
DISCORD_GUILD_ID=seu_guild_id
DISCORD_CHANNEL_ID=seu_channel_id
DB_HOST=postgres
DB_PORT=5432
DB_USER=impulso_user
DB_PASSWORD=senha_super_segura_aqui
DB_NAME=boost_tracker
MONITOR_PORT=8080
```

#### 5. Iniciar com Docker Compose
```bash
# Build das imagens
docker-compose build

# Iniciar em background
docker-compose up -d

# Ver logs
docker-compose logs -f bot

# Parar
docker-compose down
```

#### 6. Verificar Status
```bash
# Health check
curl http://localhost:8080/health

# Estatísticas
curl http://localhost:8080/stats

# Ver containers rodando
docker ps
```

---

## ✅ Opção 2: Railway.app (Mais Fácil)

### Passo a Passo:

#### 1. Criar Conta
- Acesse: https://railway.app
- Faça login com GitHub

#### 2. Criar Novo Projeto
- Clique em "New Project"
- Selecione "Deploy from GitHub repo"
- Conecte seu repositório

#### 3. Adicionar Banco de Dados PostgreSQL
```bash
# No Railway:
# - Clique em "Add"
# - Selecione "PostgreSQL"
# - Conecte ao projeto
```

#### 4. Configurar Variáveis de Ambiente
```bash
# No Railway, adicione as variáveis:
DISCORD_TOKEN=seu_token
DISCORD_GUILD_ID=seu_guild_id
DISCORD_CHANNEL_ID=seu_channel_id
DB_HOST=sua_db_host
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=boost_tracker
MONITOR_PORT=8080
```

#### 5. Deploy Automático
- Railway detecta automaticamente o `docker-compose.yml`
- Faz deploy automático a cada push no GitHub

---

## ✅ Opção 3: Heroku (Gratuito mas Limitado)

### Passo a Passo:

#### 1. Instalar Heroku CLI
```bash
# macOS
brew tap heroku/brew && brew install heroku

# Ubuntu/Linux
curl https://cli-assets.heroku.com/install-ubuntu.sh | sh

# Windows
# Download: https://cli-assets.heroku.com/heroku-x64.exe
```

#### 2. Login no Heroku
```bash
heroku login
```

#### 3. Criar App
```bash
# No diretório do projeto
heroku create seu-app-name

# Ou se já tem um repo
heroku apps:create seu-app-name
```

#### 4. Adicionar PostgreSQL
```bash
heroku addons:create heroku-postgresql:hobby-dev
```

#### 5. Configurar Variáveis
```bash
heroku config:set DISCORD_TOKEN=seu_token
heroku config:set DISCORD_GUILD_ID=seu_guild_id
heroku config:set DISCORD_CHANNEL_ID=seu_channel_id
heroku config:set MONITOR_PORT=8080

# Ver configurações
heroku config
```

#### 6. Deploy
```bash
# Fazer push para Heroku
git push heroku main

# Ver logs
heroku logs --tail
```

---

## ✅ Opção 4: Replit (Online IDE)

### Passo a Passo:

#### 1. Acessar Replit
- Vá para: https://replit.com
- Faça login/crie conta
- Clique em "Create Repl"

#### 2. Importar do GitHub
```bash
# Selecione "Import from GitHub"
# Cole: https://github.com/xre2033-hub/discord-boost-bot-automatico
```

#### 3. Instalar Dependências
```bash
# No Replit shell
go mod download
```

#### 4. Configurar Variáveis de Ambiente
```bash
# Clique em "Secrets" (cadeado)
# Adicione as variáveis:
DISCORD_TOKEN=seu_token
DISCORD_GUILD_ID=seu_guild_id
DISCORD_CHANNEL_ID=seu_channel_id
```

#### 5. Executar
```bash
# Terminal Replit
go run main.go
```

---

## ✅ Opção 5: AWS EC2 (Produção)

### Passo a Passo:

#### 1. Criar Instância EC2
- Acesse: https://console.aws.amazon.com
- Selecione: Ubuntu 22.04 LTS
- Tamanho: t2.micro (grátis por 1 ano)
- Crie keypair e baixe `.pem`

#### 2. Conectar SSH
```bash
# Mudar permissões da chave
chmod 400 sua-chave.pem

# Conectar
ssh -i sua-chave.pem ubuntu@seu-ec2-ip
```

#### 3. Instalar Docker (igual VPS)
```bash
# [Seguir passos da Opção 1]
```

#### 4. Criar RDS PostgreSQL
- No AWS console, crie uma instância RDS PostgreSQL
- Configure security groups para aceitar conexões

#### 5. Configurar Variáveis
```bash
# Editar .env com dados do RDS
DB_HOST=seu-rds-endpoint.amazonaws.com
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=sua_senha
DB_NAME=boost_tracker
```

---

## 🔧 Configuração de Domínio (Opcional)

### Para acessar via domínio customizado:

#### 1. Registrar Domínio
- Namecheap.com
- GoDaddy.com
- Google Domains

#### 2. Apontar DNS
```
# Adicionar A Record:
A  @  seu-ip-vps
```

#### 3. SSL com Let's Encrypt
```bash
sudo apt install certbot python3-certbot-nginx -y
sudo certbot certonly --standalone -d seudominio.com
```

#### 4. Configurar Nginx (Reverse Proxy)
```bash
sudo apt install nginx -y

# Configurar arquivo
sudo nano /etc/nginx/sites-available/default
```

Conteúdo:
```nginx
server {
    listen 80;
    server_name seudominio.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
    }
}
```

---

## 📊 Monitoramento em Produção

### 1. Uptime Robot (Grátis)
- https://uptimerobot.com
- Monitora seu bot a cada 5 minutos
- Avisa se cair

### 2. PM2 (Node.js/Forever Restart)
```bash
# Para manter bot sempre rodando
pm2 start "go run main.go" --name discord-bot
pm2 save
pm2 startup
```

### 3. Systemd Service
```bash
sudo nano /etc/systemd/system/discord-boost-bot.service
```

Conteúdo:
```ini
[Unit]
Description=Discord Boost Bot
After=network.target

[Service]
Type=simple
User=ubuntu
WorkingDirectory=/opt/discord-boost-bot-automatico
ExecStart=/usr/bin/docker-compose up
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
```

Ativar:
```bash
sudo systemctl daemon-reload
sudo systemctl enable discord-boost-bot
sudo systemctl start discord-boost-bot
```

---

## 💰 Comparativo de Custos

| Opção | Custo | Vantagens | Desvantagens |
|-------|-------|-----------|--------------|
| VPS (DigitalOcean) | $5-12/mês | Total controle, barato | Gerenciamento manual |
| Railway | $5-50/mês | Fácil, automático | Pode ficar caro |
| Heroku | $7+ (grátis descontinuado) | Simples | Caro para produção |
| Replit | Grátis | Online, fácil | Limitado |
| AWS | $0-50+/mês | Escalável | Complexo |

---

## 🚀 Minha Recomendação

**Para começar:** Railway.app (mais fácil)
**Para produção:** DigitalOcean + Docker

---

## ✅ Checklist de Deploy

- [ ] Repositório GitHub criado e público
- [ ] Token Discord gerado e seguro
- [ ] Guild ID e Channel ID configurados
- [ ] Banco de dados PostgreSQL rodando
- [ ] Variáveis de ambiente configuradas
- [ ] Docker e Docker Compose instalados
- [ ] Teste local: `docker-compose up`
- [ ] Teste API: `curl http://localhost:8080/health`
- [ ] Deploy em VPS/plataforma
- [ ] Monitoramento ativo (Uptime Robot)
- [ ] Logs sendo registrados no banco

---

**Precisa de ajuda com alguma etapa? Me avise!** 🎯
