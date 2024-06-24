from pydantic import BaseModel, Field
from random import uniform, randint
from datetime import datetime


class EquipmentDataModel(BaseModel):
    timestamp: datetime = Field(default_factory=datetime.now, description="Timestamp of the equipment data")
    suction_pressure: float = Field(..., description="Sucction pressure in pounds per square inch (psi)")
    discharge_pressure: float = Field(..., description="Discharge pressure in pounds per square inch (psi)")
    flow_rate: float = Field(..., description="Flow rate in gallons per minute (gpm)")
    fluid_temperature: float = Field(..., description="Fluid temperature in Celsius (C)") 
    bearing_temperature: float = Field(..., description="Bearing temperature in Celsius (C)")    
    vibration: float = Field(..., description="Vibration in millimeters per second (mm/s)")
    impeller_speed: int = Field(..., description="Impeller speed in revolutions per minute (rpm)")
    lubrication_oil_level: float = Field(..., description="Lubrication oil level in millimeters (mm)")
    npsh: float = Field(..., description="Net Positive Suction Head in meters (m)")



