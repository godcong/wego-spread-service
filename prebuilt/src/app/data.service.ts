import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class DataService {
  private client: HttpClient;

  constructor(client: HttpClient) {
    this.client = client;
  }

  getActivityList() {
    return this.client.get('http://localhost:8081/spread/activity'
    );
  }
}
