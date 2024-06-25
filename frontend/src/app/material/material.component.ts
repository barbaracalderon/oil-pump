import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { MaterialService } from '../service/material.service';

@Component({
  selector: 'app-material',
  standalone: true,
  imports: [CommonModule, HttpClientModule],
  templateUrl: './material.component.html',
  styleUrls: ['./material.component.css'],
  providers: [MaterialService]
})
export class MaterialComponent implements OnInit {
  materialData: any[] = [];

  constructor(private materialService: MaterialService) {}

  ngOnInit(): void {
    this.materialService.getMaterialData().subscribe((data: any) => {
      this.materialData = data.materials;
    });
  }

  getStatusClass(material: any): string {
    const vibration = material.vibration;
    const bearingTemperature = material.bearing_temperature;
    const impellerSpeed = material.impeller_speed;

    if (vibration < 0.1 || vibration > 5 || bearingTemperature < 20 || bearingTemperature > 100 || impellerSpeed < 1000 || impellerSpeed > 3600) {
      return 'bg-danger';
    }

    return 'bg-success';
  }
}
