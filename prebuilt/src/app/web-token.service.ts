import {Injectable} from '@angular/core';

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
