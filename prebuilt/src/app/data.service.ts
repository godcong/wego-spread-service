import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {WebTokenService} from './web-token.service';

@Injectable({
  providedIn: 'root'
})
export class DataService {
  private client: HttpClient;

  constructor(client: HttpClient) {
    this.client = client;
  }

  getActivityList() {
    return this.client.get('http://localhost:8081/spread/activity', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  getUserActivityList() {
    return this.client.get('http://localhost:8081/spread/user/activity', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

}
