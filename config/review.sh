#!/bin/bash

URL="http://localhost:3000/review"

# Arrays com frases positivas e negativas
positive_feedback=(
  "Adorei o serviço, super rápido!"
  "Muito bom, recomendo!"
  "Excelente atendimento, voltarei com certeza."
  "Ótima experiência, tudo funcionou bem."
  "Produto de qualidade, parabéns à equipe!"
)

negative_feedback=(
  "Demorou demais para responder."
  "Tive problemas e ninguém me ajudou."
  "Serviço deixou a desejar."
  "Não recomendo, experiência ruim."
  "Atendimento péssimo, fiquei decepcionado."
)

names=(
  "João" "Maria" "Carlos" "Fernanda" "Lucas"
  "Ana" "Bruno" "Juliana" "Roberto" "Cláudia"
)

for i in $(seq 1 10000); do
  # Seleciona um nome aleatório
  name="${names[$RANDOM % ${#names[@]}]}"

  # Decide aleatoriamente entre positivo ou negativo
  if (( RANDOM % 2 )); then
    desc="${positive_feedback[$RANDOM % ${#positive_feedback[@]}]}"
  else
    desc="${negative_feedback[$RANDOM % ${#negative_feedback[@]}]}"
  fi

  echo "[$i] Enviando review de $name: \"$desc\""
  curl -s -o /dev/null -w "%{http_code}\n" -X POST "$URL" \
    -H "Content-Type: application/json" \
    -d "{\"name\": \"$name\", \"description\": \"$desc\"}"
done
