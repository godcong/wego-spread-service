import {NgModule} from '@angular/core';
import {Routes, RouterModule, ExtraOptions} from '@angular/router';
import {HomeComponent} from './home/home.component';
import {MyComponent} from './my/my.component';
import {FavoriteComponent} from './favorite/favorite.component';
import {QrcodeComponent} from './qrcode/qrcode.component';
import {ActivityComponent} from './activity/activity.component';
import {MySpreadComponent} from './my/my-spread/my-spread.component';
import {MyActivityComponent} from './my/my-activity/my-activity.component';
import {MyShareComponent} from './my/my-share/my-share.component';
import {MyInfoComponent} from './my/my-info/my-info.component';

const routes: Routes = [
  // {
  // path: '',
  // component: AppComponent,
  // children: [
  {
    path: 'home', component: HomeComponent,
  },
  {
    path: 'my', component: MyComponent, children: [
      {path: 'spread', component: MySpreadComponent},
      {path: 'info', component: MyInfoComponent},
      {path: 'share', component: MyShareComponent},
      {path: 'activity', component: MyActivityComponent},
    ],
  },
  {
    path: 'favorite', component: FavoriteComponent,
  },
  {
    path: 'qrcode/:id', component: QrcodeComponent,
  },
  {
    path: 'activity/:id', component: ActivityComponent,
  },
  {
    path: '', redirectTo: '/home', pathMatch: 'full'
  }
  // ],
  // },
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
