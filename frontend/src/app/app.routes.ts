import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { MachineComponent } from './machine/machine.component';
import { FluidsComponent } from './fluids/fluids.component';
import { MaterialComponent } from './material/material.component';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'home', component: HomeComponent },
  { path: 'machine', component: MachineComponent },
  { path: 'fluids', component: FluidsComponent },
  { path: 'material', component: MaterialComponent },

];
