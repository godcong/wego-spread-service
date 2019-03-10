import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {WebTokenService} from './web-token.service';
import {ParamMap} from '@angular/router';

declare let wx: any;
const HOST = 'http://localhost:8081';

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

  getActivityList() {

    return this.client.get(HOST + '/api/v0/spread/activity', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    )
      ;
  }

  getUserActivityList() {
    return this.client.get(HOST + '/api/v0/spread/user/activity', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  getSpreadShareInfo(id: string, user: string) {
    return this.client.get(HOST + '/api/v0/spread/spread/' + id + '/share', {
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

  wxInit(params: ParamMap) {
    wx.config({
      debug: true, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
      appId: params['appID'], // 必填，公众号的唯一标识
      timestamp: params['timestamp'], // 必填，生成签名的时间戳
      nonceStr: params['nonceStr'], // 必填，生成签名的随机串
      signature: params['signature'],// 必填，签名
      jsApiList: ['onMenuShareTimeline', 'onMenuShareAppMessage', 'onMenuShareQQ', 'onMenuShareQZone'] // 必填，需要使用的JS接口列表
    });
  }

}
