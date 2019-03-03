import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class HomeDataService {
  private client: HttpClient;

  constructor(client: HttpClient) {
    this.client = client;
  }

  getActivityList() {
    this.client.get('http://localhost:8081/spread/activity').subscribe((v: any) => {
      console.log(v);
      return;
    });
    return [];
  }

}
