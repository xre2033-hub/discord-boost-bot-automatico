# ⚡ Quick Start - Instalação Rápida

## 🚀 5 Minutos para Colocar Online

### Pré-requisitos
- Docker instalado: https://docs.docker.com/get-docker/
- Discord Bot Token (veja abaixo como criar)
- Guild ID e Channel ID do Discord

---

## 📱 Passo 1: Gerar Discord Bot Token

### 1.1 Criar Aplicação no Discord Developer Portal
```
1. Acesse: https://discord.com/developers/applications
2. Clique em "New Application"
3. Dê um nome (ex: "Boost Bot")
4. Vá para "Bot" na esquerda
5. Clique em "Add Bot"
6. Em TOKEN, clique em "Copy"
```

### 1.2 Copie seu token
```
YOUR_DISCORD_TOKEN = (copie e salve em lugar seguro!)
```

### 1.3 Configurar Permissões
```
1. Em "Bot Permissions", ative:
   - Send Messages ✅
   - Embed Links ✅
   - Read Message History ✅
2. Copie a URL gerada em OAuth2 > URL Generator
3. Acesse a URL para adicionar o bot ao seu servidor
```

---

## 🆔 Passo 2: Obter Guild ID e Channel ID

### 2.1 Ativar Modo Desenvolvedor no Discord
```
Discord Settings > Advanced > Developer Mode (ON)
```

### 2.2 Obter IDs
```
1. Clique com botão direito no seu servidor > "Copy Server ID"
   YOUR_GUILD_ID = (cole aqui)

2. Clique com botão direito no canal > "Copy Channel ID"
   YOUR_CHANNEL_ID = (cole aqui)
```

---

## 📦 Passo 3: Clonar e Configurar

```bash
# 1. Clonar repositório
git clone https://github.com/xre2033-hub/discord-boost-bot-automatico.git
cd discord-boost-bot-automatico

# 2. Copiar arquivo de configuração
cp .env.example .env

# 3. Editar .env com seus dados
nano .env
```

### Conteúdo do `.env`:
```env
DISCORD_TOKEN=seu_token_aqui
DISCORD_GUILD_ID=seu_guild_id_aqui
DISCORD_CHANNEL_ID=seu_channel_id_aqui
DB_HOST=postgres
DB_PORT=5432
DB_USER=impulso_user
DB_PASSWORD=sua_senha_super_segura
DB_NAME=boost_tracker
MONITOR_PORT=8080
LOG_LEVEL=INFO
```

---

## 🐳 Passo 4: Executar com Docker

```bash
# Build das imagens
docker-compose build

# Iniciar (em background)
docker-compose up -d

# Ver logs
docker-compose logs -f bot

# Parar
docker-compose down
```

---

## ✅ Passo 5: Testar

```bash
# Health check
curl http://localhost:8080/health

# Ver estatísticas
curl http://localhost:8080/stats

# Ver histórico de boosts
curl http://localhost:8080/boosts

# Ver logs
curl http://localhost:8080/logs
```

---

## 🤖 Comandos no Discord

```
Digite no seu canal do Discord:
!status
```

Resposta esperada:
```
📊 Status do Bot de Boost
├ Total de Boosts: 0
├ Hoje: 0
├ Erros: 0
└ Status: 🟢 Online
```

---

## 🔍 Troubleshooting Rápido

| Problema | Solução |
|----------|---------|
| Token inválido | Copie novamente do Discord Developer Portal |
| Bot não aparece no servidor | Regenere OAuth2 URL e aceite as permissões |
| Banco de dados não conecta | Aguarde 10s para container PostgreSQL iniciar |
| Porta 8080 em uso | Mude `MONITOR_PORT` no .env para outra porta |
| Bot não envia mensagens | Verifique permissões do bot no canal |

---

## 📈 Próximos Passos

1. **Monitorar:** Acesse `http://seu-ip:8080/stats`
2. **Hostname:** Configure domínio customizado (veja HOSTING_GUIDE.md)
3. **Backup:** Configure backup automático do banco de dados
4. **Logs:** Analise logs em `http://seu-ip:8080/logs`

---

## 🆘 Precisa de Ajuda?

```bash
# Ver todos os logs
docker-compose logs

# Ver logs do bot apenas
docker-compose logs bot

# Ver logs do PostgreSQL
docker-compose logs postgres

# Reiniciar tudo
docker-compose restart

# Parar e remover tudo
docker-compose down -v
```

---

**🎉 Pronto! Seu bot está rodando!**

Boost será enviado **a cada hora** automaticamente! ⚡
