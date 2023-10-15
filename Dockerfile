FROM nats-streaming

# Порт NATS Streaming
EXPOSE 4222

# Запускаем NATS Streaming сервер
CMD ["nats-streaming-server"]