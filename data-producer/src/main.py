from generator.data_generator import DataGenerator
import json
import time
from datetime import datetime
from producer.data_producer import RabbitMQProducer
from config.rabbit_config import RabbitConfig

class DateTimeEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, datetime):
            return o.isoformat()
        return super().default(o)


def main():
    producer = RabbitMQProducer()
    start_time = time.time()
    try:
        while True:
            data = DataGenerator.generate_data()
            producer.send_data(data)
            time.sleep(RabbitConfig.MESSAGE_INTERVAL)

            if time.time() - start_time > RabbitConfig.RUN_TIME:
                print("Tempo de execução atingido. Encerrando o produtor...")
                break

    except KeyboardInterrupt:
        print("Shutting down producer...")
    finally:
        producer.close()


if __name__ == "__main__":
    main()