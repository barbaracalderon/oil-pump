import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { MachineComponent } from './machine/machine.component';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'home', component: HomeComponent },
  { path: 'machine', component: MachineComponent },
];
