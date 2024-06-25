// src/app/machine/machine.component.ts
import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { PressureService } from '../service/pressure.service';

@Component({
  selector: 'app-machine',
  standalone: true,
  imports: [CommonModule, HttpClientModule],
  templateUrl: './machine.component.html',
  styleUrls: ['./machine.component.css'],
  providers: [PressureService]
})
export class MachineComponent implements OnInit {
  pressureData: any[] = [];

  constructor(private pressureService: PressureService) {}

  ngOnInit(): void {
    this.pressureService.getPressureData().subscribe((data: any) => {
      this.pressureData = data.pressures;
    });
  }
}
