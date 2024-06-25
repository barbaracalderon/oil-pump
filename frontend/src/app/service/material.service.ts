import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MaterialService {
  private apiUrl = 'http://localhost:8080/material';

  constructor(private http: HttpClient) {}

  getMaterialData(): Observable<any> {
    return this.http.get<any>(this.apiUrl);
  }
}
