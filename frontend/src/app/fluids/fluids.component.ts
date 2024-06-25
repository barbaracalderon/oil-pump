import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FluidsService } from '../service/fluids.service';

@Component({
  selector: 'app-fluids',
  standalone: true,
  imports: [CommonModule, HttpClientModule],
  templateUrl: './fluids.component.html',
  styleUrls: ['./fluids.component.css'],
  providers: [FluidsService]
})
export class FluidsComponent implements OnInit {
  fluidsData: any[] = [];

  constructor(private fluidsService: FluidsService) {}

  ngOnInit(): void {
    this.fluidsService.getFluidsData().subscribe((data: any) => {
      this.fluidsData = data.fluids;
    });
  }

  getStatus(flowRate: number, fluidTemperature: number, lubricationOilLevel: number): string {
    if (flowRate < 50 || flowRate > 2000) {
      return 'Alert Flow Rate';
    }
    if (fluidTemperature < 20 || fluidTemperature > 80) {
      return 'Alert Fluid Temperature';
    }
    if (lubricationOilLevel < 10 || lubricationOilLevel > 100) {
      return 'Alert Lubrication Oil';
    }
    return 'OK';
  }
}
