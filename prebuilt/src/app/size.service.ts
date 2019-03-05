import {Injectable} from '@angular/core';

declare var window: Window;

@Injectable({
  providedIn: 'root'
})
export class SizeService {
  public height: number;
  public width: number;

  constructor() {
    this.height = window.innerHeight;
    this.width = window.innerWidth;
  }
}
