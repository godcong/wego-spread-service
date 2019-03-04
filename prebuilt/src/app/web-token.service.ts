import {Injectable} from '@angular/core';

declare var window: Window;

@Injectable({
  providedIn: 'root'
})
export class WebTokenService {

  constructor() {
  }

  static getToken() {
    return window.localStorage.getItem('token');
  }

  static setToken(token: string) {
    window.localStorage.setItem('token', token);
  }
}
