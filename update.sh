#!/bin/bash

# 🔄 Script de Atualização - Discord Boost Bot
# Atualiza o código do bot sem perder dados

echo "═══════════════════════════════════════════════════════════════"
echo "🔄 Atualizador - Discord Boost Bot"
echo "═══════════════════════════════════════════════════════════════"
echo ""

cd /opt/discord-boost-bot-automatico

# Verificar se está em um repositório git
if [ ! -d .git ]; then
    echo "❌ Erro: Não é um repositório git"
    exit 1
fi

echo "📥 Buscando atualizações..."
git fetch origin

# Verificar se há atualizações
if [ "$(git rev-parse HEAD)" == "$(git rev-parse origin/main)" ]; then
    echo "✅ Bot já está atualizado"
    exit 0
fi

echo "⚠️  Atualizações disponíveis!"
echo ""
echo "📝 Mudanças:"
git log --oneline HEAD..origin/main
echo ""

read -p "Deseja atualizar? (s/n) " -n 1 -r
echo ""

if [[ ! $REPLY =~ ^[Ss]$ ]]; then
    echo "❌ Atualização cancelada"
    exit 1
fi

echo "🛑 Parando containers..."
docker-compose down

echo "📥 Atualizando código..."
git pull origin main

echo "🏗️  Reconstruindo imagens..."
docker-compose build

echo "🚀 Reiniciando..."
docker-compose up -d

echo ""
echo "✅ Atualização concluída!"
echo "📊 Status:"
docker ps --format "table {{.Names}}\t{{.Status}}"
