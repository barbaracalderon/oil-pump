from models.data_equipment_model import EquipmentDataModel
from random import uniform, randint
from datetime import datetime

class DataGenerator:

    @staticmethod
    def generate_data():
        return EquipmentDataModel(
            timestamp=datetime.now().isoformat(),
            suction_pressure=round(uniform(0, 40), 2),  # Typical suction pressure in psi: 5-30
            discharge_pressure=round(uniform(25, 200), 2),  # Typical discharge pressure in psi: 50-150
            flow_rate=round(uniform(10, 5000), 2),  # Typical flow rate in gpm: 50-2000
            fluid_temperature=round(uniform(10, 100), 2),  # Typical fluid temperature in Celsius: 20-80
            bearing_temperature=round(uniform(10, 120), 2),  # Typical bearing temperature in Celsius: 20-100
            vibration=round(uniform(0.1, 10), 2),  # Typical vibration in mm/s: 0.1-5
            impeller_speed=randint(900, 4000),  # Typical impeller speed in rpm: 1000-3600
            lubrication_oil_level=round(uniform(1, 150), 2),  # Typical lubrication oil level in mm: 10, 100
            npsh=round(uniform(1, 20), 2)  # Typical NPSH in meters: 1-15
        )
