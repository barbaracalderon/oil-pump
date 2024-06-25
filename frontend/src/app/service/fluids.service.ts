import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FluidsService {
  private apiUrl = 'http://localhost:8080/fluid';

  constructor(private http: HttpClient) { }

  getFluidsData(): Observable<any> {
    return this.http.get<any>(this.apiUrl);
  }
}
