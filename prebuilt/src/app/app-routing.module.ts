import {NgModule} from '@angular/core';
import {Routes, RouterModule, ExtraOptions} from '@angular/router';
import {QrcodeComponent} from './qrcode/qrcode.component';
import {HomeComponent} from './home/home.component';

const routes: Routes = [
  {
    path: 'code', component: QrcodeComponent,
  },
  {
    path: 'home', component: HomeComponent,
  },
  {path: '', redirectTo: '/code', pathMatch: 'full'},
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
