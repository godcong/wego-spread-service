import {Injectable} from '@angular/core';

declare var window: Window;
const navHeight = 48;

@Injectable({
  providedIn: 'root'
})
export class SizeService {
  public height: number;
  public width: number;

  constructor() {
    this.height = window.innerHeight - navHeight;
    this.width = window.innerWidth;
  }
}
