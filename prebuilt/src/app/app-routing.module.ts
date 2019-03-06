import {NgModule} from '@angular/core';
import {Routes, RouterModule, ExtraOptions} from '@angular/router';
import {HomeComponent} from './home/home.component';
import {MyComponent} from './my/my.component';
import {FavoriteComponent} from './favorite/favorite.component';
import {QrcodeComponent} from './qrcode/qrcode.component';

const routes: Routes = [
  {
    path: 'home', component: HomeComponent,
  },
  {
    path: 'my', component: MyComponent,
  },
  {
    path: 'favorite', component: FavoriteComponent,
  },
  {
    path: 'qrcode', component: QrcodeComponent,
  },
  {
    path: '', redirectTo: '/home', pathMatch: 'full'
  },
];

const config: ExtraOptions = {
  useHash: false,
  enableTracing: true,
};

@NgModule({
  imports: [RouterModule.forRoot(routes, config)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
