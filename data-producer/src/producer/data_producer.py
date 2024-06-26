import json
import pika
import pika.credentials
from config.rabbit_config import RabbitConfig
from models.data_equipment_model import EquipmentDataModel

class RabbitMQProducer:
    def __init__(self):
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(
                host=RabbitConfig.HOST,
                port=5672,
                credentials=pika.PlainCredentials(
                    username=RabbitConfig.USER,
                    password=RabbitConfig.PASSWORD
                    )
            )
        )
        self.channel = self.connection.channel()
        self.channel.exchange_declare(
            exchange=RabbitConfig.EXCHANGE,
            exchange_type='direct',
            durable=True
        )
        self.channel.queue_declare(
            queue=RabbitConfig.QUEUE,
            durable=True
        )
        self.channel.queue_bind(
            exchange=RabbitConfig.EXCHANGE,
            queue=RabbitConfig.QUEUE,
            routing_key=RabbitConfig.ROUTING_KEY
        )


    def send_data(self, data: EquipmentDataModel):
        message = data.model_dump_json()
        self.channel.basic_publish(
            exchange=RabbitConfig.EXCHANGE,
            routing_key=RabbitConfig.ROUTING_KEY,
            body=message,
        )
        print(f"Data sent: {data}")

    def close(self):
        self.connection.close()
