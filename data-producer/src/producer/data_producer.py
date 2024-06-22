import json
import pika
import pika.credentials
from config.rabbit_config import RabbitConfig
from models.data_equipment_model import EquipmentDataModel

class RabbitMQProducer:
    def __init__(self):
        self.channel = pika.BlockingConnection(
            pika.ConnectionParameters(
            host=RabbitConfig.HOST,
            port=5672,
            credentials=pika.PlainCredentials(
                username=RabbitConfig.USER,
                password=RabbitConfig.PASSWORD
                )
            )
        ).channel()


    def send_data(self, data: EquipmentDataModel):
        message = data.model_dump_json()
        self.channel.basic_publish(
            exchange=RabbitConfig.EXCHANGE,
            routing_key='',
            body=message,
        )
        print(f"Data sent: {data}")

    def close(self):
        self.connection.close()
