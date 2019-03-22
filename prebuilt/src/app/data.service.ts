import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {WebTokenService} from './web-token.service';

const HOST = '';

interface UserActivity {
  PropertyID: string;
  ActivityID: string;
  UserID: string;
  IsStar: boolean;
  SpreadCode: string;
  IsPass: boolean;
  SpreadNumber: number;
}

interface UserActivityInterface {
  Current: number;
  Desc: boolean;
  Detail: UserActivity;
  Limit: number;
  Total: number;
  TotalPage: number;
}

@Injectable({
  providedIn: 'root'
})
export class DataService {
  private client: HttpClient;

  constructor(client: HttpClient) {
    this.client = client;
  }

  getActivityList(t: string) {
    return this.client.get(HOST + '/api/v0/spread/activity', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        },
        params: {
          type: t,
        },
      }
    );
  }

  getActivityInfo(id: string) {
    return this.client.get(HOST + '/api/v0/spread/activity/' + id, {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  getUserActivityList(favorite: boolean) {
    let url = '/api/v0/spread/user/activity/show/all';
    if (true === favorite) {
      url = '/api/v0/spread/user/activity/show/favorite';
    }
    return this.client.get(HOST + url, {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  getSpreadShareInfo(id: string, user: string) {
    return this.client.get(HOST + '/api/v0/spread/spreads/' + id + '/codes', {
        params: {
          user,
        },
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  getMyInfo() {
    return this.client.get(HOST + '/api/v0/spread/users/info', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }


  getMySpread() {
    return this.client.get(HOST + '/api/v0/spread/users/spreads', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  postActivityJoin(id: string, code: string) {
    return this.client.post(HOST + '/api/v0/spread/userActivities/' + id, null, {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        },
        params: {
          code,
        }
      }
    );
  }

  postFavoriteJoin(id: string, b: boolean) {

    return this.client.post(HOST + '/api/v0/spread/userActivities/' + id + '/favorite', null, {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        },
        params: {
          favorite: b,
        }
      }
    );
  }
}
