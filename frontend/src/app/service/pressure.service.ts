import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PressureService {
  private apiUrl = 'http://localhost:8080/pressure';

  constructor(private http: HttpClient) { }

  getPressureData(): Observable<any> {
    return this.http.get<any>(this.apiUrl);
  }
}
